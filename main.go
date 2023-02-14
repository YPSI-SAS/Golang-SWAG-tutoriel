package main

import (
	"Golang-GORM-tutoriel/controllers"
	"Golang-GORM-tutoriel/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectDb("library.db")
	database.CreateModel(db)
	database.InitDatabase(db)

	
	bookRepo := controllers.BookRepo{
		Db: db,
	}

	r := gin.Default()
	r.GET("/books", bookRepo.FindBooks)
	r.GET("/books/author/:author", bookRepo.FindBooksByAuthor)
	r.POST("/book", bookRepo.CreateBook)
	r.DELETE("/book/:id", bookRepo.DeleteBook)
	r.Run()

}
