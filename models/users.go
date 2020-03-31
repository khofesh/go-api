package models

import (
	"errors"

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
	Email        string  `bson:"email"`
	Bio          UserBio `bson:"bio"`
	PasswordHash string  `bson:"passwordHash"`
	Type         string  `bson:"type"`
	Demo         bool    `bson:"demo"`
}

func (u *UserModel) hashPassword(password string) error {
	if len(password) == 0 {
		return errors.New("Password cannot be empty")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)

	return nil
}

func (u *UserModel) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
