package models

import "github.com/gin-gonic/gin"

// UserSerializer :
type UserSerializer struct {
	c *gin.Context
}

// UserResponse : response type
type UserResponse struct {
	Email string
	Bio   UserBio
	Type  string
	Demo  bool
}
