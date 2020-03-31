package setup

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/khofesh/simple-go-api/middlewares"
	"github.com/khofesh/simple-go-api/routes"
)

// Router : initiate routers
func Router() *gin.Engine {
	// Run Gin in Release mode
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(middlewares.SecureFunc())

	r.Use(CookieSession())

	r.Use(cors.New(cors.Config{
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowOrigins:     []string{"http://localhost:3000", "https://localhost:8080"},
	}))

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// api
	v1 := r.Group("/api/v0")
	routes.UserRoute(v1.Group("/users"))

	// 404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	return r
}
