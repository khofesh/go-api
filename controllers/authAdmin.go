package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/khofesh/simple-go-api/common"
	"github.com/khofesh/simple-go-api/forms"
	"github.com/khofesh/simple-go-api/models/adminmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type loginResponse struct {
	ID    interface{} `json:"id,omitempty"`
	Email interface{} `json:"email"`
	Role  interface{} `json:"role"`
}

// AdminLogin ...
func AdminLogin(c *gin.Context) {
	var loginData forms.LoginAdminData

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	coll := common.GetCollection("simple", "admins")

	var adminData adminmodel.Model
	if err := coll.FindOne(context.TODO(), bson.M{"email": loginData.Email}).Decode(&adminData); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Wrong password or email!"})
		return
	}

	if err := adminData.CheckPassword(loginData.Password); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Wrong password or email!"})
		return
	}

	ts, err := common.CreateToken(adminData.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	session := sessions.Default(c)
	session.Options(sessions.Options{
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   60 * 60 * 24,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	session.Set("email", adminData.Email)
	session.Set("id", adminData.ID.Hex())
	session.Set("access_token", ts.AccessToken)
	session.Set("access_expire", ts.AtExpires)
	session.Set("refresh_token", ts.RefreshToken)
	session.Set("refresh_expire", ts.RtExpires)

	if err := session.Save(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":   ts.AccessToken,
		"access_expire":  ts.AtExpires,
		"refresh_token":  ts.RefreshToken,
		"refresh_expire": ts.RtExpires,
		"userData": loginResponse{
			ID:    adminData.ID.Hex(),
			Email: adminData.Email,
			Role:  adminData.Role,
		},
	})
}

// AdminLogout ...
func AdminLogout(c *gin.Context) {
	// var err error

	// _, err = common.ExtractTokenMetadata(c.Request)
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, "unauthorized")
	// 	return
	// }

	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "Successfully logged out"})
	return
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
		session := sessions.Default(c)

		_, ok := claims["refresh_uuid"].(string)
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, "Error occurred")
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

		session.Set("access_token", ts.AccessToken)
		session.Set("refresh_token", ts.RefreshToken)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		c.JSON(http.StatusCreated, tokens)
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": "refresh expired"})
	return

}

// CheckSession ...
func CheckSession(c *gin.Context) {
	session := sessions.Default(c)
	accessToken := session.Get("access_token")
	if accessToken == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    401,
			"message": "Session doesn't exist",
		})
		return
	}

	refreshToken := session.Get("refresh_token")
	userID := session.Get("id")
	email := session.Get("email")
	role := session.Get("role")
	atExpire := session.Get("access_expire")
	rtExpire := session.Get("refresh_expire")

	c.JSON(http.StatusOK, gin.H{
		"code":           200,
		"message":        "Session exists",
		"access_token":   accessToken,
		"access_expire":  atExpire,
		"refresh_token":  refreshToken,
		"refresh_expire": rtExpire,
		"userData": loginResponse{
			ID:    userID,
			Email: email,
			Role:  role,
		},
	})
}
