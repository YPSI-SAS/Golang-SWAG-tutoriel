package controllers

import (
	"Golang-GORM-tutoriel/models"
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

func (repository *BookRepo) FindBooks(c *gin.Context) {
	var bookModel models.Book
	books, err := bookModel.GetBooks(repository.Db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erreur récupération des livres")
		return
	}
	c.JSON(http.StatusOK, gin.H{"books": books})
}

func (repository *BookRepo) FindBooksByAuthor(c *gin.Context) {
	author := c.Param("author")

	var bookModel models.Book
	books, err := bookModel.GetBooksByAuthor(repository.Db, author)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erreur récupération des livres")
		return
	}
	c.JSON(http.StatusOK, gin.H{"books": books})
}

func (repository *BookRepo) CreateBook(c *gin.Context) {
	var bookInput BookCreate
	if err := c.ShouldBindJSON(&bookInput); err != nil {
		c.String(http.StatusInternalServerError, "Erreur récupération du JSON")
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

	c.JSON(http.StatusOK, gin.H{"book": newBook})
}

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

	c.JSON(http.StatusOK, gin.H{"data": true})
}
