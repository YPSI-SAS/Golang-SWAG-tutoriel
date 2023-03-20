package main

import (
	"Golang-SWAG-tutoriel/controllers"
	"Golang-SWAG-tutoriel/database"

	_ "Golang-SWAG-tutoriel/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server for demonstration.

// @contact.name   YPSI SAS
// @contact.url https://www.ypsi.fr/
// @contact.email info@ypsi.fr

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()

}
