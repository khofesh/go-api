package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khofesh/simple-go-api/common"
	"github.com/khofesh/simple-go-api/forms"
	"github.com/khofesh/simple-go-api/models/adminmodel"
)

// AdminLogin ...
func AdminLogin(c *gin.Context) {
	var user forms.SigninData

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	if user.Email != adminmodel.Example.Email || user.Password != adminmodel.Example.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	var result = adminmodel.Example

	ts, err := common.CreateToken(result.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = common.CreateAuth(result.ID, ts)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	c.JSON(http.StatusOK, tokens)
}

// AdminLogout ...
func AdminLogout(c *gin.Context) {
	var err error
	var deleted int64
	var accessDetails *common.AccessDetails

	accessDetails, err = common.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	deleted, err = common.DeleteAuth(accessDetails.AccessUUID)
	if err != nil || deleted == 0 {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	c.JSON(http.StatusOK, "Successfully logged out")
}
