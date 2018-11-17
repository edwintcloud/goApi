package members

import (
	"context"
	"fmt"
	"strings"

	"github.com/mongodb/mongo-go-driver/bson"

	"goApi/models/member"
	"goApi/utils/mongodb"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
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
		routes.PUT("", c.updateMember)
		routes.DELETE("", c.deleteMember)
	}
}

// READ ALL or READ ALL that match query params
func (*membersController) getMembers(c *gin.Context) {
	cur, err := collection.Find(context.Background(), nil)
	defer cur.Close(context.Background())
	if err == nil {
		var members []member.Member

		// build out our members slice
		for cur.Next(context.Background()) {
			member := member.Member{}
			err := cur.Decode(&member)
			if err == nil {
				q := c.Request.URL.Query()
				if len(q) > 0 { // If query params specified
					var isMatch = false
					var m = structs.Map(&member)
					for k := range q {
						if m[k] == strings.Join(q[k], "") {
							isMatch = true
						} else {
							isMatch = false
							break
						}
					}
					if isMatch {
						members = append(members, member)
					}
				} else {
					members = append(members, member)
				}
			}
		}

		if len(members) > 0 {
			c.JSON(200, members)
			return
		}
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

// UPDATE ANY MATCHING QUERY PARAMS
func (*membersController) updateMember(c *gin.Context) {
	member := member.Member{}
	if c.ShouldBind(&member) == nil {
		q := c.Request.URL.Query()
		if len(q) > 0 { // If query params specified
			var filter = make(map[string]interface{})
			for k := range q {
				filter[k] = strings.Join(q[k], "")
			}
			reqBody := structs.Map(&member)
			for k := range reqBody {
				if len(reqBody[k].(string)) < 3 {
					delete(reqBody, k)
				}
			}
			update := bson.M{"$set": reqBody}
			if res, err := collection.UpdateMany(context.Background(), filter, update); err == nil && res.ModifiedCount != 0 {
				c.JSON(200, gin.H{
					"message": fmt.Sprintf("Successfull updated %d members!", res.ModifiedCount),
				})
				return
			}
		}
		c.JSON(400, gin.H{
			"error": "No members updated!",
		})
		return
	}
	c.JSON(400, gin.H{
		"error": "Invalid input!",
	})
}

// DELETE ONE
func (*membersController) deleteMember(c *gin.Context) {
	c.JSON(400, gin.H{
		"error": "Not Implemented!",
	})
}
