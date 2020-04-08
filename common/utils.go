package common

import (
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandString : generate random string
func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// NBSecretPassword : secret password
const NBSecretPassword = "A String Very Very Very Strong!!@##$!@#$"

// NBRandomPassword : random password
const NBRandomPassword = "A String Very Very Very Niubilty!!@##$!@#4"

// GenToken : generate  jwt token (used in the request header)
func GenToken(id uint) string {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, _ := jwtToken.SignedString([]byte(NBSecretPassword))
	return token
}

// Bind :
func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}
