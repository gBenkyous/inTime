package main

import (
	"github.com/gin-gonic/gin"
	"intimeServer/api/handler"
)

// @title gin-swagger todos
// @version 1.0
// @license.name intime
// @description 見本apiです
func main() {
	r := gin.Default()
	handler.SetHandlers(r)
	r.Run(":8080")
}
