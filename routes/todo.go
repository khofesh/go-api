package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khofesh/simple-go-api/controllers"
	"github.com/khofesh/simple-go-api/middlewares"
)

// TodoRoute ...
func TodoRoute(router *gin.RouterGroup) {
	router.POST("/create", middlewares.TokenAuthMiddleware(), controllers.CreateTodo)
}
