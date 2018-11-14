package db

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson/objectid"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func Connect() {
	client, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// test connection by inserting a document
	collection := client.Database("goApi").Collection("test")
	res, err := collection.InsertOne(context.Background(), map[string]string{"hello": "world"})
	if err != nil {
		log.Fatal(err)
	}

	// output result of inserted document
	if oid, ok := res.InsertedID.(objectid.ObjectID); ok {
		fmt.Println("Id inserted: ", oid.Hex())
	}
}
