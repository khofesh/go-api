package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// UserBio : user's biography
type UserBio struct {
	firstname  string `bson:"firstname"`
	middlename string `bson:"middlename"`
	lastname   string `bson:"lastname"`
	gender     string `bson:"gender"`
	picture    string `bson:"picture"`
}

// UserModel : user's model
type UserModel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email      string             `bson:"email" json:"email" binding:"required"`
	Bio        UserBio            `bson:"bio" json:"bio" binding:"omitempty"`
	Password   string             `bson:"password" json:"password" binding:"required"`
	Type       string             `bson:"type" json:"type" binding:"required"`
	Demo       bool               `bson:"demo" json:"demo" binding:"required"`
	Securities []SecuritiesModel  `bson:"securities,omitempty" json:"securities"`
}

// HashPassword : generate password hash
func (u *UserModel) HashPassword(password string) error {
	if len(password) == 0 {
		return errors.New("Password cannot be empty")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(passwordHash)

	return nil
}

// CheckPassword : check if login password matches hashed password
func (u *UserModel) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
