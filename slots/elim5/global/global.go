package global

import (
	"elim5/utils/timer"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"gorm.io/gorm/logger"
	"strings"
	"sync"

	"golang.org/x/sync/singleflight"

	"elim5/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB      *gorm.DB // 读写库
	GVA_READ_DB *gorm.DB // 只读库 用于查询 不可用于写操作
	NOLOG_DB    *gorm.DB // 没有日志的读写库
	GVA_DBList  map[string]*gorm.DB
	GVA_REDIS   *redis.ClusterClient
	GVA_CONFIG  config.Server
	GVA_VP      *viper.Viper
	// GVA_LOG    *oplogging.Logger
	GVA_LOG                 *ZapLogger
	GVA_Timer               timer.Timer = timer.NewTimerTask()
	GVA_Concurrency_Control             = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
	SvName     string

	Json = sonic.ConfigFastest
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GVA_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GVA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}

func NoLog(tx *gorm.DB) *gorm.DB {
	return tx.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)})
}

func GetListenUrl(port string) string {
	if !strings.Contains(port, ":") {
		return fmt.Sprintf("%s:%s", GVA_CONFIG.System.ListenIp, port)
	}
	return port
}

func GetConnectUrl(port string) string {
	if !strings.Contains(port, ":") {
		return fmt.Sprintf("%s:%s", GVA_CONFIG.System.ConnectIp, port)
	}
	return port
}

func GetLogLevel() logger.LogLevel {
	switch GVA_CONFIG.Mysql.GetLogMode() {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	default:
		return logger.Info
	}
}
