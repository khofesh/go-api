package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khofesh/simple-go-api/controllers"
)

// SignupRoute : handle route for signing up
func SignupRoute(router *gin.RouterGroup) {
	router.POST("/signup", controllers.Signup)
}
