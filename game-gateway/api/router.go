package api

import (
	"fmt"
	"game-gateway/service"

	"github.com/gin-gonic/gin"
)

func NewApi() *gin.Engine {
	engine := gin.Default()

	echoRouter := engine.Group("/ktpd")
	{

		echoRouter.GET("/room/:grop", service.RoomGameConn)

		echoRouter.GET("/hello", func(context *gin.Context) {
			fmt.Println("world")
		})

	}

	return engine
}
