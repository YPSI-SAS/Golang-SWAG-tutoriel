package controllers

import (
	"Golang-SWAG-tutoriel/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookRepo struct {
	Db *gorm.DB
}

type BookCreate struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BooksList struct {
	Books []models.Book `json:"books"`
}

// @Summary Get all books
// @Description Get all books in DB
// @Tags book
// @Success 200 {object} controllers.BooksList "Books return"
// @Failure 500 "Error to find"
// @Router /books [get]
func (repository *BookRepo) FindBooks(c *gin.Context) {
	var bookModel models.Book
	books, err := bookModel.GetBooks(repository.Db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erreur récupération des livres")
		return
	}
	c.JSON(http.StatusOK, BooksList{Books: *books})
}

// @Summary Get all books by author name
// @Description Get all books in DB by author name
// @Tags book
// @Success 200 {object} controllers.BooksList "Books return"
// @Failure 500 "Error to find book by author"
// @Failure 400 "Error on request"
// @Param author path string true "Author's name"
// @Router /books/author/{author} [get]
func (repository *BookRepo) FindBooksByAuthor(c *gin.Context) {
	author := c.Param("author")

	var bookModel models.Book
	books, err := bookModel.GetBooksByAuthor(repository.Db, author)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erreur récupération des livres")
		return
	}
	c.JSON(http.StatusOK, BooksList{Books: *books})
}

// @Summary Create book
// @Description Create book in DB
// @Tags book
// @Success 200 {object} models.Book "Books return"
// @Failure 500 "Error to create"
// @Failure 400 "Error on request"
// @Param body body controllers.BookCreate true "Body"
// @Router /book [post]
func (repository *BookRepo) CreateBook(c *gin.Context) {
	var bookInput BookCreate
	if err := c.ShouldBindJSON(&bookInput); err != nil {
		c.String(http.StatusBadRequest, "Erreur récupération du JSON")
		return
	}

	newBook := models.Book{
		Title:  bookInput.Title,
		Author: bookInput.Author,
	}

	err := newBook.UpdateOrCreateBook(repository.Db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erreur création du livre")
		return
	}

	c.JSON(http.StatusOK, newBook)
}

// @Summary Delete one book
// @Description Delete one book in DB by ID
// @Tags book
// @Success 200
// @Failure 500 "Error to delete"
// @Failure 400 "Error on request"
// @Failure 404 "Error book not find"
// @Param id path string true "Book's ID"
// @Router /book/{id} [delete]
func (repository *BookRepo) DeleteBook(c *gin.Context) {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Erreur récupération du paramètre")
		return
	}

	var bookFind models.Book
	err = bookFind.GetBookById(repository.Db, uint(bookId))
	if err != nil {
		c.String(http.StatusNotFound, "Le livre n'existe pas")
		return
	}

	err = bookFind.DeleteBook(repository.Db, uint(bookFind.ID))
	if err != nil {
		c.String(http.StatusInternalServerError, "Erreur suppression du livre")
		return
	}

	c.Status(http.StatusOK)
}
