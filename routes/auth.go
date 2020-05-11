package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khofesh/simple-go-api/controllers"
	"github.com/khofesh/simple-go-api/middlewares"
)

// AuthRoute : handle route for signing up
func AuthRoute(router *gin.RouterGroup) {
	router.POST("/signup", controllers.Signup)
	router.POST("/login", middlewares.JWTMiddleware.LoginHandler)
	router.GET("/refresh-token", middlewares.JWTMiddleware.RefreshHandler)
	router.GET("/logout", middlewares.JWTMiddleware.LogoutHandler)
}
