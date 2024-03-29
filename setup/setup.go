package setup

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khofesh/simple-go-api/middlewares"
	"github.com/khofesh/simple-go-api/routes"
)

// Router : initiate routers
func Router() *gin.Engine {
	// Run Gin in Release mode
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.2", "127.0.0.1"})
	r.Use(middlewares.SecureFunc())

	middlewares.InitJWT()

	r.Use(RedisSession())

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// api routes
	v0 := r.Group("/api/v0")
	routes.UserRoute(v0.Group("/users"))
	routes.AuthRoute(v0.Group("/auth"))
	routes.TodoRoute(v0.Group("/todo"))

	v0Admin := r.Group("/api/v0/admin")
	routes.AuthAdminRoute(v0Admin)
	v0Admin.Use(middlewares.TokenAuthMiddleware())
	routes.CRUDAdminRoutes(v0Admin.Group("/crud"))

	// 404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	return r
}
