package routerType

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "intimeServer/docs"
)

// MakeBookHandlers make url handlers
func MakeSwaggerHandlers(r *gin.Engine) {

	//	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	//http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
