package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRetrieve :
// get user's data
func UserRetrieve(c *gin.Context) {
	session := sessions.Default(c)
	email := session.Get("user_email")

	c.JSON(http.StatusOK, gin.H{"message": "Success getting user", "email": email})
}
