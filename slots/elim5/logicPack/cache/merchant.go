package cache

import (
	"context"
	"elim5/enum"
	"elim5/global"
	"elim5/model/business"
	"elim5/model/common"
	"elim5/utils/helper"
	"errors"
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"sync/atomic"
)

var (
	merchantCacheById    = cmap.NewWithCustomShardingFunction[uint, *business.Merchant](sharding)
	merchantCacheByAgent = cmap.New[*business.Merchant]()
	merchantCacheOk      = atomic.Bool{}
)

func checkMerchantCache() error {
	if merchantCacheOk.Load() {
		return nil
	}
	var all []*business.Merchant
	err := global.GVA_READ_DB.Omit("appkey", "remark").Find(&all).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			merchantCacheOk.Store(true)
			return nil
		}
		global.GVA_LOG.Error("获取商户列表失败", zap.Error(err))
		return enum.ErrCommon
	}
	for _, merchant := range all {
		// 测试商户设置
		if merchant.ID == 1 {
			if global.GVA_CONFIG.System.TestApiUrl != "" {
				merchant.ApiUrl = global.GVA_CONFIG.System.TestApiUrl
			}
		}

		merchantCacheById.Set(merchant.ID, merchant)
		merchantCacheByAgent.Set(merchant.Agent, merchant)
	}
	merchantCacheOk.Store(true)
	return nil
}

func GetMerchant(anyKey any) (m *business.Merchant, err error) {
	err = checkMerchantCache()
	if err != nil {
		return nil, err
	}

	var ok bool
	switch k := anyKey.(type) {
	case string:
		m, ok = merchantCacheByAgent.Get(k)
	case uint:
		m, ok = merchantCacheById.Get(k)
	case int:
		m, ok = merchantCacheById.Get(uint(k))
	case int32:
		m, ok = merchantCacheById.Get(uint(k))
	}
	if !ok {
		return nil, fmt.Errorf("merchant: %v not found", anyKey)
	}

	return m, err
}

func GetMerchantMap() map[uint]*business.Merchant {
	_ = checkMerchantCache()
	return merchantCacheById.Items()
}

func ClearMerchant() {
	merchantCacheOk.Store(false)
	merchantCacheById.Clear()
	merchantCacheByAgent.Clear()
	merchantCacheOk.Store(false)
}

const MerchantExtraParamsKey = "{merchant_extra_params}"

func GetMerchantExtraParamsKey(merchantId uint) string {
	return strconv.Itoa(int(merchantId))
}

func GetMerchantExtraParams(merchantId uint) (params common.MerchantExtraParams, err error) {
	key := GetMerchantExtraParamsKey(merchantId)
	var jsonStr string
	jsonStr, err = global.GVA_REDIS.HGet(context.Background(), MerchantExtraParamsKey, key).Result()
	if err == nil {
		_ = global.Json.UnmarshalFromString(jsonStr, &params)
		return
	}

	err = global.GVA_READ_DB.Model(&business.Merchant{}).Where("id = ?", merchantId).First(&params).Error
	if err != nil {
		global.GVA_LOG.Error("GetMerchantExtraParams err: " + err.Error())
		return
	}
	err = SetMerchantExtraParamsKey(merchantId, params)
	if err != nil {
		global.GVA_LOG.Error("SetMerchantExtraParamsKey err: " + err.Error())
	}
	return
}

func SetMerchantExtraParamsKey(merchantId uint, params common.MerchantExtraParams) (err error) {
	s, _ := global.Json.MarshalToString(params)
	return global.GVA_REDIS.HSet(context.Background(), MerchantExtraParamsKey, GetMerchantExtraParamsKey(merchantId), s).Err()
}

func DelMerchantExtraParams(merchantIds ...uint) (err error) {
	return global.GVA_REDIS.HDel(context.Background(), MerchantExtraParamsKey, helper.StringArr(merchantIds)...).Err()
}

func ClearMerchantExtraParams() (err error) {
	return global.GVA_REDIS.Del(context.Background(), MerchantExtraParamsKey).Err()
}
