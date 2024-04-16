package cache

import (
	"elim5/enum"
	"elim5/global"
	"elim5/model/business"
	cmap "github.com/orcaman/concurrent-map/v2"
	"sync/atomic"
)

var (
	ipWhiteListCache   = cmap.New[struct{}]()
	ipWhiteListCacheOk = atomic.Bool{}
)

func checkIpWhiteListCache() error {
	if ipWhiteListCacheOk.Load() {
		return nil
	}
	var ips []string
	err := global.GVA_READ_DB.
		Model(&business.IpWhite{}).
		Where("status = ?", enum.Yes).
		Pluck("ip", &ips).Error
	if err != nil {
		return err
	}
	for _, ip := range ips {
		ipWhiteListCache.Set(ip, struct{}{})
	}
	ipWhiteListCacheOk.Store(true)
	return nil
}

func GetIPWhiteListMap() (m map[string]struct{}, err error) {
	err = checkIpWhiteListCache()
	if err != nil {
		return nil, err
	}
	return ipWhiteListCache.Items(), nil
}

func AddIPWhiteList(ip string) {
	ipWhiteListCache.Set(ip, struct{}{})
}

func ClearIPWhiteListCache() {
	ipWhiteListCacheOk.Store(false)
	ipWhiteListCache.Clear()
	ipWhiteListCacheOk.Store(false)
}
