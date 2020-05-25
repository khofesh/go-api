package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khofesh/simple-go-api/controllers"
	"github.com/khofesh/simple-go-api/middlewares"
)

// AuthAdminRoute ...
func AuthAdminRoute(router *gin.RouterGroup) {
	router.POST("/login", controllers.AdminLogin)
	router.POST("/refresh-token", controllers.RefreshToken)
	router.GET("/check-session", controllers.CheckSession)
	router.GET("/logout", middlewares.TokenAuthMiddleware(), controllers.AdminLogout)
}

// CRUDAdminRoutes ...
func CRUDAdminRoutes(router *gin.RouterGroup) {
	router.POST("/create-admin", controllers.CreateOneAdmin)
}
