package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khofesh/simple-go-api/controllers"
)

// AuthRoute : handle route for signing up
func AuthRoute(router *gin.RouterGroup) {
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)
}
