package main

import (
	"goApi/controllers/api"
	"goApi/controllers/users"
	"log"
	"os"

	"goApi/utils/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set gin to production mode
	gin.SetMode(gin.DebugMode)

	// connect to mysql
	db.Connect(os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DB"))

	// close database when this function block gets the SIG_INT
	defer db.Close()

	// create new instance of gin
	r := gin.New()

	// set router to use default middlewares
	r = gin.Default()

	// initialize controllers
	api.Init(r)
	users.Init(r)

	// start the server
	r.Run(":5000")
}
