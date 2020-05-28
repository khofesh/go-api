package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khofesh/simple-go-api/common"
)

// Todo ...
type Todo struct {
	UserID string `json:"user_id"`
	Title  string `json:"title"`
}

// CreateTodo ...
func CreateTodo(c *gin.Context) {
	var td *Todo
	var userID string

	if err := c.ShouldBindJSON(&td); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	extractedToken, err := common.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	userID, err = common.CompareData(c, extractedToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	td.UserID = userID
	c.JSON(http.StatusCreated, td)
}
