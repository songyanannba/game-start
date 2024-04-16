package core

import (
	"elim5/global"
	"elim5/initialize"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func BaseInit() {
	global.GVA_VP = Viper() // 初始化Viper

	global.GVA_LOG = Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG.Logger)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.DBList()
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}
}
