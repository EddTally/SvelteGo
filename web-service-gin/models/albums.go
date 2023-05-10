package models

import (
	"gorm.io/gorm"
)

type Albums struct {
	gorm.Model
	ID     uint    `gorm:"primary key;autoIncrement" json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func MigrateAlbums(db *gorm.DB) error {
	//Drop table
	db.Migrator().DropTable(&Albums{})
	// Recreate it
	err := db.AutoMigrate(&Albums{})

	// Now add the default data
	db.Create(&Albums{Title: "Blue Train", Artist: "John Coltrane", Price: 56.99})
	db.Create(&Albums{Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99})
	db.Create(&Albums{Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99})

	return err
}
