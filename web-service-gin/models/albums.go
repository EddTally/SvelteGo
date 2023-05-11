package models

import (
	"time"

	"gorm.io/gorm"
)

type Album struct {
	ID        uint `gorm:"primary key;autoIncrement" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Title     string         `json:"title"`
	Artist    string         `json:"artist"`
	Price     float64        `json:"price"`
}

func MigrateAlbums(db *gorm.DB) error {
	//Drop table
	db.Migrator().DropTable(&Album{})
	// Recreate it
	err := db.AutoMigrate(&Album{})

	// Now add the default data
	db.Create(&Album{Title: "Blue Train", Artist: "John Coltrane", Price: 56.99})
	db.Create(&Album{Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99})
	db.Create(&Album{Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99})

	return err
}
