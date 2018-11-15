package members

import (
	"context"
	"fmt"
	"goApi/models/member"
	"goApi/utils/mongodb"

	"github.com/mongodb/mongo-go-driver/bson/objectid"

	"github.com/gin-gonic/gin"
)

type membersController struct{}

// Init initializes our controllers and routes
func Init(e *gin.Engine) {
	c := membersController{}

	// routes
	routes := e.Group("/members")
	{
		routes.GET("", c.getMembers)
		routes.POST("", c.createMember)
	}
}

func (*membersController) getMembers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Not implemented",
	})
}

func (*membersController) createMember(c *gin.Context) {
	collection := mongodb.Client.Collection("Members")
	member := member.Member{}

	if vErr := member.CheckValid(&member); c.ShouldBind(&member) == nil && vErr == nil {
		res, err := collection.InsertOne(context.Background(), member.HashPassword(&member))
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Unable to create new member!",
			})
		} else {
			id := res.InsertedID.(objectid.ObjectID).Hex()
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("New member with id: %s inserted into database!", id),
			})
		}
	} else {
		c.JSON(400, gin.H{
			"error": vErr.Error(),
		})
	}
}
