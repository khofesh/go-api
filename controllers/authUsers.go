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
	var data forms.LoginUserData

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

	type ResponseData struct {
		Email string         `json:"email"`
		Bio   models.UserBio `json:"bio"`
		Type  string         `json:"type"`
		Demo  bool           `json:"demo"`
		Token string         `json:"token"`
	}

	var sessionData = ResponseData{
		Email: result.Email,
		Bio:   result.Bio,
		Type:  result.Type,
		Demo:  result.Demo,
		Token: common.GenToken(result.ID),
	}

	var userData = ResponseData{
		Email: result.Email,
		Bio:   result.Bio,
		Type:  result.Type,
		Demo:  result.Demo,
	}

	session := sessions.Default(c)
	session.Options(sessions.Options{
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   60 * 60,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	session.Set("user_token", sessionData.Token)
	session.Set("user_email", sessionData.Email)
	session.Set("user_id", result.ID.String())

	if session.Save() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in!", "userData": userData})
}

// Logout : handle user's log out
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out!"})
}

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
