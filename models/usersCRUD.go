package models

import (
	"context"
	"fmt"
	"log"

	"github.com/khofesh/simple-go-api/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateUser : create user
func CreateUser(data interface{}) {
	collection := common.GetCollection("simple", "users")

	idxMod := mongo.IndexModel{
		Keys: bson.M{"username": 1}, Options: options.Index().SetUnique(true),
	}

	names, idxerr := collection.Indexes().CreateOne(context.TODO(), idxMod)
	if idxerr != nil {
		log.Fatal(idxerr)
	}
	fmt.Printf("created indexes %v\n", names)

	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// FindOneUser : find a user
func FindOneUser() {}

// UpdateUser : update user's data
func UpdateUser() {}

// DeleteUser : delete user's data
func DeleteUser() {}
