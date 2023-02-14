package database

import (
	"Golang-GORM-tutoriel/models"
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)


func ConnectDb(host string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(host), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect database")
	}
	return db
}

func CreateModel(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&models.Book{})
	return db
}

func InitDatabase(db *gorm.DB) {
	var book1 = models.Book{
		Title:  "Voyage au centre de la Terre",
		Author: "Jules Verne",
	}
	book1.UpdateOrCreateBook(db)

	var book2 = models.Book{
		Title:  "Vingt Mille Lieues sous les mers",
		Author: "Jules Verne",
	}
	book2.UpdateOrCreateBook(db)

	var book3 = models.Book{
		Title:  "Les Misérables",
		Author: "Victor Hugo",
	}
	book3.UpdateOrCreateBook(db)
}
