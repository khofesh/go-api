package models

import (
	"context"
	"errors"
	"fmt"

	"github.com/khofesh/simple-go-api/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateUser : create user
func (u *UserModel) CreateUser() error {
	collection := common.GetCollection("simple", "users")

	if val, _ := collection.CountDocuments(context.TODO(), bson.M{"email": u.Email}); val != 0 {
		return errors.New("Email already exists")
	}

	idxMod := mongo.IndexModel{
		Keys: bson.M{"email": 1}, Options: options.Index().SetUnique(true),
	}

	names, err := collection.Indexes().CreateOne(context.TODO(), idxMod)
	if err != nil {
		return err
	}
	fmt.Printf("created indexes %v\n", names)

	_, err = collection.InsertOne(context.TODO(), u)
	if err != nil {
		return err
	}

	return nil
}

// FindOneUser : find a user
func FindOneUser(filter bson.M) (UserModel, error) {
	var result UserModel

	coll := common.GetCollection("simple", "users")

	if err := coll.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

// UpdateUser : update user's data
func (u *UserModel) UpdateUser(update bson.M) error {
	coll := common.GetCollection("simple", "users")

	_, err := coll.UpdateOne(context.TODO(), bson.M{"email": u.Email}, update)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser : delete user's data
func DeleteUser(filter bson.M) error {
	coll := common.GetCollection("simple", "users")

	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}
