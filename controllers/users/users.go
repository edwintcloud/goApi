package users

import (
	"fmt"
	"goApi/models/user"
	"goApi/utils/db"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type usersController struct{}

// Init initializes our controller and routes and runs migrations for controller
func Init(e *gin.Engine) {

	// run migrations
	db.DB.AutoMigrate(&user.User{})

	c := usersController{}

	// routes
	routes := e.Group("/users")
	{
		routes.GET("", c.getUsers)
		routes.POST("", c.createUser)
		routes.PUT("", c.updateUser)
		routes.DELETE("", c.deleteUser)

		routes.POST("/login", c.loginUser)
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

	if c.ShouldBind(&user) == nil && user.Username != "" && len(user.Password) > 3 {
		db.DB.Create(&user)
		c.JSON(200, user)
	} else {
		c.JSON(400, gin.H{
			"message": "Unable to create user",
		})
	}
}

// UPDATE ONE by body or query
func (*usersController) updateUser(c *gin.Context) {
	user := user.User{}

	if c.ShouldBind(&user) == nil && user.ID != 0 {
		db.DB.Model(&user).Updates(user)
		c.JSON(200, user)
	} else if c.ShouldBindQuery(&user) == nil && user.ID != 0 {
		db.DB.Model(&user).Updates(user)
		c.JSON(200, user)
	} else {
		c.JSON(400, gin.H{
			"message": "Unable to update any users",
		})
	}
}

// DELETE ONE by query string
func (*usersController) deleteUser(c *gin.Context) {
	user := user.User{}

	if c.ShouldBindQuery(&user) == nil {
		// Make sure our primary key has a value so we don't wipe the table
		if user.ID == 0 {
			c.JSON(400, gin.H{
				"message": "Invalid query, no users deleted.",
			})
		} else {
			db.DB.Delete(&user)
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("ID %d deleted from database", user.ID),
			})
		}
	}
}

// Login user with req body
func (*usersController) loginUser(c *gin.Context) {
	user, foundUser := user.User{}, user.User{}

	if c.ShouldBind(&user) == nil {
		if err := db.DB.Where("Username = ?", user.Username).First(&foundUser).Error; err == nil {
			if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err == nil {
				c.JSON(200, gin.H{
					"message": "You are now logged in!",
				})
			} else {
				c.JSON(400, gin.H{
					"message": "Password incorrect",
				})
			}
		} else {
			c.JSON(400, gin.H{
				"message": "User not found!",
			})
		}
	}
}
