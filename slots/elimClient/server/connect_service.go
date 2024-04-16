package server

import (
	"github.com/gin-gonic/gin"
)

func GameConn(ctx *gin.Context) {
	GameClientManager.Register(ctx)
}
