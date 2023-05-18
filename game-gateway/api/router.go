package api

import (
	"fmt"
	"game-gateway/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewApi() *gin.Engine {
	engine := gin.Default()

	echoRouter := engine.Group("/ktpd")
	{

		echoRouter.GET("/room", service.RoomGameConn)

		echoRouter.GET("/hello", func(context *gin.Context) {
			fmt.Println("world")
		})

		echoRouter.GET("/hello1", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello, World!",
			})
		})

	}

	return engine
}
