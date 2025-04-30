package routes

import (
	"zocket-task/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	router.POST("/usersignup", authHandler.UserSignup)
	router.POST("/userlogin", authHandler.UserLogin)
}
