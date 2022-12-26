package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RoomRegister(ctx *gin.Context) {
	userId := ctx.GetHeader("userId")
	if len(userId) <= 0 {
		fmt.Println("缺少userId")
		return
	}

	//up := websocket.

}
