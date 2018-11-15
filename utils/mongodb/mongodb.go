package mongodb

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// Client is the currently connected mongo database session
var Client *mongo.Database

// Connect connect to database
func Connect(nDb string) {
	res, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	err = res.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	Client = res.Database(nDb)
}
