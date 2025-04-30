package routes

import (
	"zocket-task/pkg/api/handlers"
	"zocket-task/pkg/utils/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, authHandler *handlers.AuthHandler, taskHandler *handlers.TaskHandler) {
	router.POST("/usersignup", authHandler.UserSignup)
	router.POST("/userlogin", authHandler.UserLogin)
	authGroup := router.Group("/task")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.POST("/createTask", taskHandler.CreateTask)
		authGroup.GET("/trackTask", taskHandler.TrackTask)
		authGroup.POST("/assignTask", taskHandler.AssignTask)
		authGroup.POST("/suggestTask", taskHandler.GetTaskSuggestions)
	}
}
