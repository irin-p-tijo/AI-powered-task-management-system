package api

import (
	"log"
	"zocket-task/pkg/api/handlers"
	"zocket-task/pkg/api/routes"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(
	authHandler *handlers.AuthHandler,
) *ServerHTTP {

	router := gin.New()

	router.Use(gin.Logger())

	/////////////////routes//////////////////////////

	routes.UserRoutes(router.Group("/user"), authHandler)
	return &ServerHTTP{engine: router}
}

func (sh *ServerHTTP) Start(infoLog *log.Logger, errorLog *log.Logger) {

	infoLog.Printf("starting server on :8000")
	err := sh.engine.Run(":8000")
	errorLog.Fatal(err)
}
