package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khofesh/simple-go-api/models"
)

// Login : handle user's login
func Login() {}

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
