package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khofesh/simple-go-api/common"
)

// TokenAuthMiddleware ...
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := common.ValidateToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
