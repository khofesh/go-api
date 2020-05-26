package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khofesh/simple-go-api/controllers"
	"github.com/khofesh/simple-go-api/middlewares"
)

// UserRoute user related routes
func UserRoute(router *gin.RouterGroup) {
	router.GET("/test", middlewares.JWT.MiddlewareFunc(), controllers.UserRetrieve)
}
