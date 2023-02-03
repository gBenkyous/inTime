package routerType

import (
	"github.com/gin-gonic/gin"
	"intimeServer/usecase/user"
)

// MakeBookHandlers make url handlers
func MakeUserHandlers(r *gin.Engine) {

	r.GET("/user", user.GetUser)

}
