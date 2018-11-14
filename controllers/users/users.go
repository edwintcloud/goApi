package users

import (
	"fmt"
	"goApi/models/user"
	"goApi/utils/db"
	"log"

	"github.com/gin-gonic/gin"
)

type usersController struct{}

// Init initializes our controller and routes
func Init(e *gin.Engine) {
	c := usersController{}

	// routes
	routes := e.Group("/users")
	{
		routes.GET("", c.getUsers)
		routes.POST("", c.createUser)
		routes.PUT("", c.updateUser)
		routes.DELETE("", c.deleteUser)
	}

}

// READ ALL
func (*usersController) getUsers(c *gin.Context) {
	var users []user.User
	if err := db.DB.Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
		log.Fatal(err)
	} else {
		c.JSON(200, users)
	}
}

// CREATE ONE
func (*usersController) createUser(c *gin.Context) {
	user := user.User{}

	if c.ShouldBind(&user) == nil {
		db.DB.Create(&user)
		c.JSON(200, user)
	}
}

// UPDATE ONE by query string
func (*usersController) updateUser(c *gin.Context) {
	user := user.User{}

	if c.ShouldBindQuery(&user) == nil {
		db.DB.Model(&user).Updates(user)
		c.JSON(200, user)
	}
}

// DELETE ONE by query string
func (*usersController) deleteUser(c *gin.Context) {
	user := user.User{}

	if c.ShouldBindQuery(&user) == nil {
		db.DB.Delete(&user)
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("ID %d deleted from database", user.ID),
		})
	}
}
