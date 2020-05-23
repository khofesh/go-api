package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khofesh/simple-go-api/common"
	"github.com/khofesh/simple-go-api/models/adminmodel"
)

// CreateOneAdmin ...
func CreateOneAdmin(c *gin.Context) {

	// check authentication
	tokenAuth, err := common.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	_, err = common.FetchAuth(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	var adminData adminmodel.Model

	if err = c.ShouldBindJSON(&adminData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	if err = adminData.HashPassword(adminData.Password); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
	}

	if err = adminData.CreateAdmin(); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin account is successfully created!"})
}
