package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/khofesh/simple-go-api/common"
	"github.com/khofesh/simple-go-api/forms"
	"github.com/khofesh/simple-go-api/models/adminmodel"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// RefreshToken ...
func RefreshToken(c *gin.Context) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	refreshToken := mapToken["refresh_token"]

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_REFRESH_KEY")), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, "Refresh token expired")
		return
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUUID, ok := claims["refresh_uuid"].(string)
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, "Error occurred")
			return
		}

		deleted, delErr := common.DeleteAuth(refreshUUID)
		if delErr != nil || deleted == 0 {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}

		// convert userID (string) to ObjectID
		objectID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}

		ts, createErr := common.CreateToken(objectID)
		if createErr != nil {
			c.JSON(http.StatusForbidden, createErr.Error())
			return
		}

		err = common.CreateAuth(objectID, ts)
		if err != nil {
			c.JSON(http.StatusForbidden, err.Error())
			return
		}

		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		c.JSON(http.StatusCreated, tokens)
		return
	}

	c.JSON(http.StatusUnauthorized, "refresh expired")
	return

}
