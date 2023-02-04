package user

import (
	"github.com/gin-gonic/gin"
	"intimeServer/internal/model"
)

// GetUser ...
/*
@Summary Userをjsonで返す
@Tags User
@Produce  json
@Success 200 {object} model.User
@Router /user [get]
*/
func GetUser(c *gin.Context) {
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
