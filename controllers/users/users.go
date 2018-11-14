package users

import (
	"fmt"
	"goApi/models/user"
	"goApi/utils/db"

	"github.com/gin-gonic/gin"
)

type usersController struct{}

// Init initializes our controller and routes
func Init(e *gin.Engine) {
	c := usersController{}

	// routes
	e.GET("/users", c.getUsers)
}

func (*usersController) getUsers(c *gin.Context) {
	var users []user.User
	if err := db.DB.Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Print(err)
	} else {
		c.JSON(200, users)
	}
}
