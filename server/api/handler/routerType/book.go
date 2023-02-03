package routerType

import (
	"github.com/gin-gonic/gin"
	"intimeServer/usecase/book"
)

//MakeBookHandlers make url handlers
func MakeBookHandlers(r *gin.Engine) {

	r.GET("/books", book.GetBook)
	r.GET("/books/:id", book.GetBookByID)
	/* 下記をコマンドラインに打ち込むことでテストできる
	curl http://127.0.0.1:8080/books \
	--include \
	--header "Content-Type: application/json" \
	--request "POST" \
	--data '{"ID": "4", "Title": "testHONMA", "Author": "J.HONMA"}'
	*/
	r.POST("/books", book.PostBook)

}
