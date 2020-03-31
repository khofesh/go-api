package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khofesh/simple-go-api/controllers"
)

// UserRoute user related routes
func UserRoute(router *gin.RouterGroup) {
	router.GET("/test", controllers.UserRetrieve)
}
