package api

import "github.com/gin-gonic/gin"

type ApiController struct{}

func Init(e *gin.Engine) {
	a := ApiController{}

	// routes
	e.GET("/", a.getIndex)
}

func (r *ApiController) getIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to my API 🎉!",
	})
}
