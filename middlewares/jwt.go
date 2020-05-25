package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"
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

	coll := common.GetCollection("simple", "users")

	type ResponseData struct {
		ID    primitive.ObjectID `json:"id,omitempty"`
		Email string             `json:"email"`
		Bio   models.UserBio     `json:"bio"`
		Type  string             `json:"type"`
		Demo  bool               `json:"demo"`
	}

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "simpleApi",
		SigningAlgorithm: "HS256",
		Key:              []byte(os.Getenv("JWT_ACCESS_KEY")),
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
			var data forms.LoginUserData

			if err := c.ShouldBindJSON(&data); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			var result models.UserModel
			if err := coll.FindOne(context.TODO(), bson.M{"email": data.Email}).Decode(&result); err != nil {
				return "", jwt.ErrFailedAuthentication
			}

			if err := result.CheckPassword(data.Password); err != nil {
				return "", jwt.ErrFailedAuthentication
			}

			c.Set("useremail", result.Email)

			var userData = &ResponseData{
				ID:    result.ID,
				Email: result.Email,
				Bio:   result.Bio,
				Type:  result.Type,
				Demo:  result.Demo,
			}

			return userData, nil
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
		LoginResponse: func(c *gin.Context, code int, token string, t time.Time) {
			type userData struct {
				ID    interface{} `json:"id,omitempty"`
				Email interface{} `json:"email"`
				Type  interface{} `json:"type"`
				Demo  interface{} `json:"demo"`
			}

			email := c.MustGet("useremail")

			var result models.UserModel
			if err := coll.FindOne(context.TODO(), bson.M{"email": email}).Decode(&result); err != nil {
				c.JSON(http.StatusForbidden, gin.H{"message": jwt.ErrFailedAuthentication})
				return
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
			session.Set("user_email", result.Email)
			session.Set("user_id", result.ID.Hex())
			session.Set("user_type", result.Type)
			session.Set("user_demo", result.Demo)
			session.Set("user_token", token)

			if err := session.Save(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(code, gin.H{
				"code":   code,
				"token":  token,
				"expire": t,
				"userData": userData{
					Email: result.Email,
					ID:    result.ID.Hex(),
					Type:  result.Type,
					Demo:  result.Demo,
				},
			})
		},
		LogoutResponse: func(c *gin.Context, code int) {
			session := sessions.Default(c)
			session.Clear()
			session.Save()

			c.JSON(code, gin.H{"code": code, "message": "Successfully logged out!"})
		},
		TokenLookup:    "header: Authorization, query: token, cookie: token",
		TokenHeadName:  "Bearer",
		TimeFunc:       time.Now,
		SendCookie:     false,
		SecureCookie:   false, //non HTTPS dev environments
		CookieHTTPOnly: true,  // JS can't modify
		CookieDomain:   "localhost",
		CookieName:     "token", // default jwt
	})

	JWTMiddleware = authMiddleware

	return authMiddleware, err
}
