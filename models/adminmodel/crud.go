package adminmodel

import (
	"context"
	"errors"
	"fmt"

	"github.com/khofesh/simple-go-api/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword ...
func (u *Model) HashPassword(password string) error {
	if len(password) == 0 {
		return errors.New("Password cannot be empty")
	}

	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(passwordHash)

	return nil
}

// CheckPassword ...
func (u *Model) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

// CreateAdmin ...
func (u *Model) CreateAdmin() error {
	coll := common.GetCollection("simple", "admins")

	if val, _ := coll.CountDocuments(context.TODO(), bson.M{"email": u.Email}); val != 0 {
		return errors.New("Email already exists")
	}

	idxMod := mongo.IndexModel{
		Keys: bson.M{"email": 1}, Options: options.Index().SetUnique(true),
	}

	names, err := coll.Indexes().CreateOne(context.TODO(), idxMod)
	if err != nil {
		return err
	}
	fmt.Printf("created indexes %v\n", names)

	_, err = coll.InsertOne(context.TODO(), u)
	if err != nil {
		return err
	}

	return nil
}

// UpdateAdmin ...
func (u *Model) UpdateAdmin(update bson.M) error {
	coll := common.GetCollection("simple", "admins")

	_, err := coll.UpdateOne(context.TODO(), bson.M{"email": u.Email}, update)
	if err != nil {
		return err
	}

	return nil
}

// DeleteAdmin ...
func DeleteAdmin(filter bson.M) error {
	coll := common.GetCollection("simple", "admins")

	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

// FindOneAdmin ...
func FindOneAdmin(filter bson.M) (Model, error) {
	var result Model

	coll := common.GetCollection("simple", "admins")

	if err := coll.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}
