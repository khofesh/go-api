package models

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khofesh/simple-go-api/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateUser : create user
func (u *UserModel) CreateUser(c *gin.Context) error {
	collection := common.GetCollection("simple", "users")

	if val, _ := collection.CountDocuments(context.TODO(), bson.M{"email": u.Email}); val != 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Email already exists"})
	}

	idxMod := mongo.IndexModel{
		Keys: bson.M{"email": 1}, Options: options.Index().SetUnique(true),
	}

	names, idxerr := collection.Indexes().CreateOne(context.TODO(), idxMod)
	if idxerr != nil {
		return idxerr
	}
	fmt.Printf("created indexes %v\n", names)

	_, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		return err
	}

	return nil
}

// FindOneUser : find a user
func FindOneUser() {}

// UpdateUser : update user's data
func UpdateUser() {}

// DeleteUser : delete user's data
func DeleteUser() {}
