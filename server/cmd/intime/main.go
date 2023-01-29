package main

import (
	"fmt"
	"io"
	//"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"intimeServer/pkg/dbtest"
	//"intimeServer/pkg/keisan"
)

type book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

type HttpResponse struct {
	Message     string
	Status      int
	Description string
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	// var i int = 1
	// i = keisan.Tasizan(1, 2)
	// fmt.Println(i)
	// log.Println("OK")
	logFile, err := os.Create("production.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	//	router := gin.Default()
	router.Use(gin.CustomRecovery(ErrorHandler))
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
	//http://転送されるIP:8080/books で開く
	router.Run(":8080")
}

func getBook(c *gin.Context) {
	var err1 = 0
	fmt.Println(err1)
	var a = 1 / err1
	fmt.Println(a)
	c.IndentedJSON(http.StatusOK, books)
}

func postBook(c *gin.Context) {
	var newBook book

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

func ErrorHandler(c *gin.Context, err any) {
	goErr := errors.Wrap(err, 2)
	httpResponse := HttpResponse{Message: "Internal server error", Status: 500, Description: goErr.Error()}
	c.AbortWithStatusJSON(500, httpResponse)
}
