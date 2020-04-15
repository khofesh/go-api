package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/khofesh/simple-go-api/common"
	"github.com/khofesh/simple-go-api/forms"
	"github.com/khofesh/simple-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// JWTMiddleware ...
var JWTMiddleware *jwt.GinJWTMiddleware

// InitJWT : create auth middleware
func InitJWT() (*jwt.GinJWTMiddleware, error) {
	var identityKey = "id"

	type ResponseData struct {
		ID    primitive.ObjectID `json:"id,omitempty"`
		Email string             `json:"email"`
		Bio   models.UserBio     `json:"bio"`
		Type  string             `json:"type"`
		Demo  bool               `json:"demo"`
		Token string             `json:"token"`
	}

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "simpleApi",
		SigningAlgorithm: "HS256",
		Key:              []byte(common.NBSecretPassword),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		IdentityKey:      identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*ResponseData); ok {
				fmt.Println(v)
				return jwt.MapClaims{
					identityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &ResponseData{
				ID: claims[identityKey].(primitive.ObjectID),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var data forms.SigninData

			err := c.ShouldBindJSON(&data)
			if err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			coll := common.GetCollection("simple", "users")

			var result models.UserModel
			if err = coll.FindOne(context.TODO(), bson.M{"email": data.Email}).Decode(&result); err != nil {
				return "", jwt.ErrFailedAuthentication
			}

			err = result.CheckPassword(data.Password)
			if err != nil {
				return "", jwt.ErrFailedAuthentication
			}

			var userData = &ResponseData{
				ID:    result.ID,
				Email: result.Email,
				Bio:   result.Bio,
				Type:  result.Type,
				Demo:  result.Demo,
			}

			session := sessions.Default(c)
			session.Options(sessions.Options{
				Path:     "/",
				Domain:   "localhost",
				MaxAge:   60 * 60,
				Secure:   false,
				HttpOnly: true,
				SameSite: http.SameSiteStrictMode,
			})
			session.Set("user_email", userData.Email)
			session.Set("user_id", result.ID.String())

			if session.Save() != nil {
				return "", err
			}

			return userData, err
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			session := sessions.Default(c)
			userEmail := session.Get("user_email")

			if v, ok := data.(*models.UserModel); ok && v.Email == userEmail {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	JWTMiddleware = authMiddleware

	return authMiddleware, err
}
