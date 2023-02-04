package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "intimeServer/docs"
	"intimeServer/internal/model"
)

// インターフェースアダプタ層
var books = []model.Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

// GetBooks ...
/*
@Summary bookinfoをjsonで返す
@Tags Book
@Produce  json
@Success 200 {object} []model.Book
@Router /books [get]
*/
func GetBook(c *gin.Context) {
	c.JSON(200, books)
}

// 本の登録
/*
@Summary book情報登録
@Tags Book
@Description  本の情報を登録します
@Accept  json
@Produce  json
@Param data body model.Book true "book data"
@Success 200 {object} model.Book
@Router /books [post]
*/
func PostBook(c *gin.Context) {
	var newBook model.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// 127.0.0.1:8080/books/3 この形式
// GetBooks ...
/*
@Summary bookinfoをIDで検索しjsonで返す
@Tags Book
@Produce  json
@Param id path int true "book ID"
@Success 200 {object} model.Book
@Router /books/{id} [get]
*/
func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
