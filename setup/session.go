package setup

import (
	"fmt"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// CookieSession : setup gin-gonic sessions
func CookieSession() gin.HandlerFunc {
	var sessionSecret string = os.Getenv("SESSION_SECRET")

	store := cookie.NewStore([]byte(sessionSecret))
	return sessions.Sessions("simple_session", store)
}

// RedisSession : setup redis session
func RedisSession() gin.HandlerFunc {
	var sessionSecret string = os.Getenv("SESSION_SECRET")

	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte(sessionSecret))
	if err != nil {
		fmt.Println(err.Error())
	}
	return sessions.Sessions("simple_session", store)
}
