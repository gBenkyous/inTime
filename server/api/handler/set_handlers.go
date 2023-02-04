package handler

import (
	"github.com/gin-gonic/gin"
	"intimeServer/api/handler/router"
	"intimeServer/internal/model"
	"intimeServer/pkg/dbtest"
	"net/http"
)

var books = []model.Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

func SetHandlers(r *gin.Engine) {

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code": "404", "message": "Page not found",
		})
	})

	// https://medium.com/since-i-want-to-start-blog-that-looks-like-men-do/%E5%88%9D%E5%BF%83%E8%80%85%E3%81%AB%E9%80%81%E3%82%8A%E3%81%9F%E3%81%84interface%E3%81%AE%E4%BD%BF%E3%81%84%E6%96%B9-golang-48eba361c3b4
	//book
	router.MakeBookHandlers(r)
	router.MakeUserHandlers(r)
	router.MakeSwaggerHandlers(r)
	r.GET("/dbtest", getDbTest)

}
func getDbTest(c *gin.Context) {
	dbtest.DbTest()
	c.IndentedJSON(http.StatusOK, books)
}
