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
	router.GET("/logout", middlewares.TokenAuthMiddleware(), controllers.AdminLogout)
}
