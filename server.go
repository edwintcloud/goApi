package main

import (
	"goApi/controllers/api"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set gin to production mode
	gin.SetMode(gin.DebugMode)

	r := gin.New()

	// set router to use default middlewares
	r = gin.Default()

	// initialize api controller
	api.Init(r)

	// start the server
	r.Run(":5000")
}
