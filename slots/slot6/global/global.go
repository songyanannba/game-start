package global

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/songzhibin97/gkit/cache/singleflight"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"slot6/config"
	"slot6/utils/timer"
	"sync"
)

var (
	GVA_DB     *gorm.DB
	GVA_DBList map[string]*gorm.DB
	GVA_REDIS  *redis.Client

	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
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

type ZapLogger struct {
	*zap.Logger
}

func (z ZapLogger) Skip(s int) *ZapLogger {
	z.Logger = z.WithOptions(zap.AddCallerSkip(s))
	return &z
}

func (z ZapLogger) Fatal(v ...interface{}) {
	z.WithOptions(zap.AddCallerSkip(1)).Fatal(fmt.Sprint(v...))
}

func (z ZapLogger) Fatalf(format string, v ...interface{}) {
	z.WithOptions(zap.AddCallerSkip(1)).Fatal(fmt.Sprintf(format, v...))
}

func (z ZapLogger) Println(v ...interface{}) {
	z.WithOptions(zap.AddCallerSkip(1)).Info(fmt.Sprint(v...))
}

func (z ZapLogger) Infof(format string, a ...any) {
	z.WithOptions(zap.AddCallerSkip(1)).Info(fmt.Sprintf(format, a...))
}

func GetClusterUrl(port string) string {
	return fmt.Sprintf("%s:%s", GVA_CONFIG.System.ListenIp, port)
}
