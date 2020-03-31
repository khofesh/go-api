package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRetrieve :
// get user's data
func UserRetrieve(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Success getting user"})
}
