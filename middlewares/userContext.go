package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/khofesh/simple-go-api/common"
	"github.com/khofesh/simple-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

// UpdateContextUserModel : A helper to write user_id and user_model to the context
func UpdateContextUserModel(c *gin.Context, myUserID uint) {
	var myUserModel models.UserModel
	if myUserID != 0 {
		coll := common.GetCollection("simple", "users")
		coll.FindOne(context.TODO(), bson.M{"_id": myUserID}).Decode(&myUserModel)
	}
	c.Set("myUserID", myUserID)
	c.Set("myUserModel", myUserModel)
}
