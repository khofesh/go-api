package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/khofesh/simple-go-api/common"
	"github.com/khofesh/simple-go-api/forms"
	"github.com/khofesh/simple-go-api/models/adminmodel"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TokenDetails ...
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

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

	ts, err := CreateToken(result.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = CreateAuth(result.ID, ts)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	c.JSON(http.StatusOK, tokens)
}

// CreateToken ...
func CreateToken(userid primitive.ObjectID) (*TokenDetails, error) {
	var secretAccess string = os.Getenv("JWT_ACCESS_KEY")
	var secretRefresh string = os.Getenv("JWT_REFRESH_KEY")

	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUUID = uuid.Must(uuid.NewRandom()).String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.Must(uuid.NewRandom()).String()

	var err error

	// access token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = userid.Hex()
	atClaims["exp"] = td.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(secretAccess))
	if err != nil {
		return nil, err
	}

	// refresh token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = userid.Hex()
	rtClaims["exp"] = td.RtExpires

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(secretRefresh))
	if err != nil {
		return nil, err
	}

	return td, nil
}

// CreateAuth ...
func CreateAuth(userid primitive.ObjectID, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	client := common.GetRedis()

	err := client.Set(td.AccessUUID, userid.Hex(), at.Sub(now)).Err()
	if err != nil {
		return err
	}

	err = client.Set(td.RefreshUUID, userid.Hex(), rt.Sub(now)).Err()
	if err != nil {
		return err
	}

	return nil
}
