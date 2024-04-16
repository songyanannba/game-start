package cache

import (
	"context"
	"elim5/global"
	"elim5/utils"
	"strconv"
	"time"
)

func GetTemplateGenID(id int) string {
	return "{template_gen}:" + strconv.Itoa(id)
}

// GetTemplateGenCache 获取模版生成缓存
func GetTemplateGenCache(genId int) bool {
	key := GetTemplateGenID(genId)
	_, err := global.GVA_REDIS.Get(context.Background(), key).Bytes()
	if err != nil {
		return false
	}
	return true
}

// SetTemplateGenCache 设置模版生成缓存
func SetTemplateGenCache(genId int) (err error) {
	key := GetTemplateGenID(genId)
	var data []byte
	err = global.GVA_REDIS.SetNX(context.Background(), key, data, 5*time.Hour).Err()
	return
}

func DeleteTemplateGenCache(genId int) error {
	err := utils.DelCache(GetTemplateGenID(genId))
	if err != nil {
		return err
	}
	return nil
}
