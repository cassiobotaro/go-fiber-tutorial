package book

import (
	"net/http"

	"github.com/cassiobotaro/go-fiber-tutorial/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// Book model
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

// GetBooks ...
func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

// GetBook ...
func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

// NewBook ...
func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(http.StatusUnprocessableEntity).Send(err)
		return
	}
	db.Create(&book)
	c.JSON(book)
}

// DeleteBook ...
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(http.StatusNotFound).Send("No book found with given id")
		return
	}
	db.Delete(&book)
	c.Status(http.StatusNoContent).Send()
}
