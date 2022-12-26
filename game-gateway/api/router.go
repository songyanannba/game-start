package api

import (
	"fmt"
	"game-gateway/connect"
	"github.com/gin-gonic/gin"
)

func NewApi() *gin.Engine {
	engine := gin.Default()

	echoRouter := engine.Group("/ktpd")
	{
		echoRouter.GET("/hello" , func(context *gin.Context) {
			fmt.Println("world")
		})
		echoRouter.POST("room:channel" ,connect.RoomGameConn)


	}

	return engine
}