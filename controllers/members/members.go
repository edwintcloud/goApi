package members

import (
	"context"
	"fmt"
	"goApi/models/member"
	"goApi/utils/mongodb"

	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/mongodb/mongo-go-driver/bson/objectid"

	"github.com/gin-gonic/gin"
)

type membersController struct{}

var collection *mongo.Collection

// Init initializes our controllers and routes
func Init(e *gin.Engine) {
	c := membersController{}

	// set collection
	collection = mongodb.Client.Collection("Members")
	// routes
	routes := e.Group("/members")
	{
		routes.GET("", c.getMembers)
		routes.POST("", c.createMember)
	}
}

// READ ALL
func (*membersController) getMembers(c *gin.Context) {
	cur, err := collection.Find(context.Background(), nil)
	defer cur.Close(context.Background())
	if err == nil {
		var members []member.Member
		for cur.Next(context.Background()) {
			member := member.Member{}
			err := cur.Decode(&member)
			if err == nil {
				members = append(members, member)
			}
		}
		c.JSON(200, members)
		return
	}
	c.JSON(400, gin.H{
		"error": "Unable to find members!",
	})
}

// CREATE ONE
func (*membersController) createMember(c *gin.Context) {
	member := member.Member{}

	if c.ShouldBind(&member) == nil {
		if err := member.CheckValid(&member); err == nil {
			res, err := collection.InsertOne(context.Background(), member.HashPassword(&member))
			if err == nil {
				id := res.InsertedID.(objectid.ObjectID).Hex()
				c.JSON(200, gin.H{
					"message": fmt.Sprintf("New member with id: %s inserted into database!", id),
				})
				return
			}
		} else {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
	c.JSON(400, gin.H{
		"error": "Unable to create member!",
	})
}
