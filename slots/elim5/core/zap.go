package core

import (
	"elim5/core/internal"
	"elim5/global"
	"elim5/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap() (logger *global.ZapLogger) {
	if ok, _ := utils.PathExists(global.GVA_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.GVA_CONFIG.Zap.Director)
		_ = os.Mkdir(global.GVA_CONFIG.Zap.Director, os.ModePerm)
	}
	logger = new(global.ZapLogger)
	cores := internal.Zap.GetZapCores()
	zapLogger := zap.New(zapcore.NewTee(cores...))

	if global.GVA_CONFIG.Zap.ShowLine {
		zapLogger = zapLogger.WithOptions(zap.AddCaller())
	}
	logger.Logger = zapLogger
	return logger
}
