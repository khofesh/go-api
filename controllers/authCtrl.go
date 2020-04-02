package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignupData :
type SignupData struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Type     string `json:"type" binding:"required"`
	Demo     string `json:"demo" binding:"required"`
}

// Login : handle user's login
func Login() {}

// Logout : handle user's log out
func Logout() {}

// Signup : handle signing up
func Signup(c *gin.Context) {
	var data SignupData

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	return
}
