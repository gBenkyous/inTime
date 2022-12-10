package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "intimeServer/docs"
	"intimeServer/model"
	"intimeServer/pkg/dbtest"
	"intimeServer/pkg/keisan"
)

var books = []model.Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

// @title gin-swagger todos
// @version 1.0
// @license.name intime
// @description 見本apiです
func main() {
	var i int = 1
	i = keisan.Tasizan(1, 2)
	fmt.Println(i)
	log.Println("OK")

	router := gin.Default()
	router.GET("/books", getBook)
	// 127.0.0.1:8080/books/3
	router.GET("/books/:id", getBookByID)
	router.GET("/dbtest", getDbTest)
	/* 下記をコマンドラインに打ち込むことでテストできる
		curl http://127.0.0.1:8080/books \
	    --include \
	    --header "Content-Type: application/json" \
	    --request "POST" \
	    --data '{"ID": "4", "Title": "testHONMA", "Author": "J.HONMA"}'
	*/
	router.POST("/books", postBook)
	//	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//http://転送されるIP:8080/books で開く
	router.Run(":8080")
}

// GetBooks ...
// @Summary bookinfoをjsonで返す
// @Tags Book
// @Accept  json
// @Produce  json
// @Success 200 {object} model.User
// @Router /books [get]
func getBook(c *gin.Context) {
	user := model.User{
		ID:        13218,
		Name:      "User1",
		Lastname:  "Ivanov",
		Email:     "user@mail.com",
		IsAdmin:   false,
		AuthCount: 159,
	}

	c.JSON(200, user)
}

type ser struct {
	ID        int    `json:"id"`
	Name      string `json:"name" example:"Jack"`
	Lastname  string `json:"lastname" example:"Packard"`
	Email     string `json:"email" example:"test@mail.com"`
	IsAdmin   bool   `json:"isAdmin"`
	AuthCount int    `json:"authCount"`
}

type ResponseHello struct {
	SayHello  string
	Number    int
	Something interface{}
}

func postBook(c *gin.Context) {
	var newBook model.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func getDbTest(c *gin.Context) {
	dbtest.DbTest()
	c.IndentedJSON(http.StatusOK, books)
}
