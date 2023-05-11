package main

import (
	"example/web-service-gin/models"
	"example/web-service-gin/storage"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// Alot of the code is based upon this tutorial:
// https://github.com/AkhilSharma90/go-postgres-gorm/blob/master/main.go

// https://go.dev/doc/tutorial/web-service-gin
// album represents data about a record album.
// type album struct {
// 	ID     uint    `gorm:"primary key;autoIncrement" json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// Our Repo
type Repository struct {
	DB *gorm.DB
}

func main() {

	// Loading env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database")
	}
	err = models.MigrateAlbums(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}
	// Now R is our new repo
	r := Repository{
		DB: db,
	}

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(logFormatter))
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"PUT", "GET", "OPTIONS", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Accept", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Hello": "World"})
	})
	router.GET("/albums", r.getAlbums)
	router.GET("/albums/:id", r.getAlbumByID)
	router.POST("/albums", r.createAlbum)

	//router.Use(cors.Default())
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func (r *Repository) getAlbums(c *gin.Context) {
	albumModels := &[]models.Album{}

	err := r.DB.Find(albumModels).Error
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "album fetched successfully",
		"data":    albumModels,
	})
}

// createAlbum adds an album from JSON received in the request body.
func (r *Repository) createAlbum(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Checking that the title doesn't already exist, don't need yet
	fmt.Printf("This is the new album before adding: %+v\n", newAlbum)

	// Add the new album to the database.
	//albums = append(albums, newAlbum)
	err := r.DB.Create(&newAlbum).Error
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Could not create album",
		})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Album Created Successfully",
		"data":    newAlbum,
	})
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (r *Repository) getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	albumModel := &models.Album{}

	if id == "" {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "id cannot be empty",
		})
		return
	}

	err := r.DB.Where("id = ?", id).First(albumModel).Error
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "book fetched successfully",
		"data":    albumModel,
	})
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	// for _, a := range albums {
	// 	if a.ID == id {
	// 		c.IndentedJSON(http.StatusOK, a)
	// 		return
	// 	}
	// }
}

// Custom formatter to be passed in as m iddleware
func logFormatter(param gin.LogFormatterParams) string {

	// your custom format
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n Request: %v\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
		param.Request,
	)
}
