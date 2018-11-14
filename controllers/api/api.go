package api

import "github.com/gin-gonic/gin"

type apiController struct{}

// Init initializes our controller and routes
func Init(e *gin.Engine) {
	c := apiController{}

	// routes
	e.GET("/", c.getIndex)
}

func (*apiController) getIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to my API 🎉!",
	})
}
