package setup

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// CookieSession : setup gin-gonic sessions
func CookieSession() gin.HandlerFunc {
	var sessionSecret string = os.Getenv("SESSION_SECRET")

	store := cookie.NewStore([]byte(sessionSecret))
	return sessions.Sessions("simple_session", store)
}
