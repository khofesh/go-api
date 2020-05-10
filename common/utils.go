package common

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
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

// AccessDetails ...
type AccessDetails struct {
	AccessUUID string
	UserID     string
}

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
func GenToken(id primitive.ObjectID) string {
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

	client := GetRedis()

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

// ExtractToken ...
func ExtractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	strArr := strings.Split(token, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// VerifyToken ...
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_ACCESS_KEY")), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

// ValidateToken ...
func ValidateToken(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// ExtractTokenMetadata ...
func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			return nil, err
		}

		return &AccessDetails{
			AccessUUID: accessUUID,
			UserID:     userID,
		}, nil
	}

	return nil, err
}

// FetchAuth : lookup token metadata in redis
func FetchAuth(authD *AccessDetails) (string, error) {
	client := GetRedis()
	userID, err := client.Get(authD.AccessUUID).Result()
	if err != nil {
		return "", err
	}

	return userID, nil
}

// DeleteAuth ...
func DeleteAuth(givenUUID string) (int64, error) {
	client := GetRedis()

	deleted, err := client.Del(givenUUID).Result()
	if err != nil {
		return 0, err
	}

	return deleted, nil
}
