package cache

import (
	"elim5/enum"
	"elim5/global"
	"elim5/model/business"
	"github.com/songzhibin97/gkit/cache/local_cache"
)

const BlackListCacheKey = "black_list"

// GetCountryBlackList 获取黑名单
func GetCountryBlackList() map[string]struct{} {
	res, exist := Cache.Get(BlackListCacheKey)
	if exist && res != nil {
		return res.(map[string]struct{})
	}
	var list []string
	err := global.GVA_READ_DB.Model(&business.BlackList{}).
		Where("status = ?", enum.Yes).
		Pluck("country", &list).Error
	if err != nil {
		global.GVA_LOG.Error("GetBlackList err: " + err.Error())
		return nil
	}
	var m = make(map[string]struct{})
	for _, country := range list {
		m[country] = struct{}{}
	}
	Cache.Set(BlackListCacheKey, m, local_cache.NoExpire)
	return m
}
