package core

import (
	"go.uber.org/zap"
	"slot6/global"
	"slot6/initialize"
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
