package service

import (
	"github.com/gin-gonic/gin"
)

func RoomGameConn(ctx *gin.Context) {
	GameClientManager.RoomRegister(ctx)
}
