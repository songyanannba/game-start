package cache

// uint sharding
var sharding = func(key uint) uint32 {
	return uint32(key)
}

// ClearGameCache 清除游戏内存级配置
func ClearGameCache() {
	// 清除游戏配置
	ClearSlotCache()
	// 清除全局配置
	Cache.Flush()
}

// ClearMerchantCache 清除商户内存级相关配置
func ClearMerchantCache() {
	ClearMerchant()
	ClearIPWhiteListCache()
}
