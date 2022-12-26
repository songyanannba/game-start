package connect

import (
	"game-gateway/service"
	"github.com/gin-gonic/gin"
)

func RoomGameConn(ctx *gin.Context)  {
	service.RoomRegister(ctx)
}
