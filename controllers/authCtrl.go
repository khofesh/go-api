package controllers

import (
	"context"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/khofesh/simple-go-api/common"
	"github.com/khofesh/simple-go-api/forms"
	"github.com/khofesh/simple-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Login : handle user's login
func Login(c *gin.Context) {
	var data forms.SigninData

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	coll := common.GetCollection("simple", "users")

	var result models.UserModel
	if err = coll.FindOne(context.TODO(), bson.M{"email": data.Email}).Decode(&result); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Wrong password or email!"})
		return
	}

	err = result.CheckPassword(data.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Wrong password or email!"})
		return
	}

	type ReturnData struct {
		Email string         `json:"email"`
		Bio   models.UserBio `json:"bio"`
		Type  string         `json:"type"`
		Demo  bool           `json:"demo"`
		Token string         `json:"token"`
	}

	var sessionData = ReturnData{
		Email: result.Email,
		Bio:   result.Bio,
		Type:  result.Type,
		Demo:  result.Demo,
		Token: common.GenToken(result.ID),
	}

	var userData = ReturnData{
		Email: result.Email,
		Bio:   result.Bio,
		Type:  result.Type,
		Demo:  result.Demo,
	}

	session := sessions.Default(c)
	session.Set("sessionData", sessionData)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in!", "userData": userData})
}

// Logout : handle user's log out
func Logout() {}

// Signup : handle signing up
func Signup(c *gin.Context) {
	var data models.UserModel

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
	}

	if err = data.HashPassword(data.Password); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
	}

	if err = data.CreateUser(); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "User is successfully created!"})
}
