package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRetrieve :
// get user's data
func UserRetrieve(c *gin.Context) {
	session := sessions.Default(c)
	token := session.Get("user_token")
	fmt.Println(c.Request)
	c.JSON(http.StatusOK, gin.H{"message": "Success getting user", "token": token})
}
