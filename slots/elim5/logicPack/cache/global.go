package cache

import (
	"elim5/enum"
	"elim5/global"
	"elim5/model/business"
	"elim5/pbs/common"
	"elim5/utils"
	"elim5/utils/helper"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"github.com/songzhibin97/gkit/cache/local_cache"
)

var Cache = local_cache.NewCache()

const BumperConfigCacheKey = "bumper"

const BumperMerchantMapKey = "bumper_merchant"

// GetBumperConfigMapByCache 获取保险杠配置 全局通用
func GetBumperConfigMapByCache() map[string]*business.SlotBumper {
	res, exist := Cache.Get(BumperConfigCacheKey)
	if exist {
		return res.(map[string]*business.SlotBumper)
	}
	var bumpers []*business.SlotBumper
	global.GVA_DB.Find(&bumpers)

	m := make(map[string]*business.SlotBumper)
	merchantMap := make(map[uint]struct{})
	for _, bumper := range bumpers {
		key := bumper.Currency
		if bumper.MerchantId != 0 {
			key += "_" + helper.Itoa(bumper.MerchantId)
			merchantMap[bumper.MerchantId] = struct{}{}
		}
		m[key] = bumper
	}
	Cache.Set(BumperConfigCacheKey, m, local_cache.NoExpire)
	Cache.Set(BumperMerchantMapKey, merchantMap, local_cache.NoExpire)
	return m
}

func MerchantInBumperConfig(merchantId uint) bool {
	m, exist := Cache.Get(BumperMerchantMapKey)
	if exist {
		_, exist = m.(map[uint]struct{})[merchantId]
		return exist
	}
	return false
}

// GetBumperConfig 获取单条保险杠配置
func GetBumperConfig(currency string, merchantId uint) *business.SlotBumper {
	m := GetBumperConfigMapByCache()
	if len(m) == 0 {
		return nil
	}
	bumper, ok := m[currency+"_"+helper.Itoa(merchantId)]
	if !ok {
		bumper = m[currency]
	}
	return bumper
}

const AmountLimitConfigCacheKey = "amount_limit"

// GetAmountLimitByCache 获取余额限制 全局通用
func GetAmountLimitByCache(agent, currency string) int64 {
	res, exist := Cache.Get(AmountLimitConfigCacheKey)
	if exist && res != nil {
		return res.(map[string]int64)[agent+"_"+currency]
	}
	var bumpers []*business.SlotAmountLimit
	global.GVA_DB.Find(&bumpers)
	m := lo.SliceToMap(bumpers, func(item *business.SlotAmountLimit) (string, int64) {
		return item.Agent + "_" + item.Currency, item.Limit
	})
	Cache.Set(AmountLimitConfigCacheKey, m, local_cache.NoExpire)
	return m[agent+"_"+currency]
}

const FileLocalCacheKey = "files"

// GetFileMapByCache 获取文件列表至本地缓存
func GetFileMapByCache() (map[uint][]*common.File, error) {
	res, exist := Cache.Get(FileLocalCacheKey)
	if exist && res != nil {
		return res.(map[uint][]*common.File), nil
	}

	files, err := GetFileList()
	if err != nil {
		return nil, enum.ErrNotOpen
	}

	m := make(map[uint][]*common.File)
	for _, file := range files {
		m[file.GameId] = append(m[file.GameId], &common.File{
			Name: file.Name,
			Path: file.Path,
		})
	}
	Cache.Set(FileLocalCacheKey, m, local_cache.NoExpire)
	return m, nil
}

func GetPublicFileListByCache() ([]*common.File, error) {
	fileMap, err := GetFileMapByCache()
	if err != nil {
		return nil, err
	}
	if list, ok := fileMap[0]; ok {
		return list, nil
	}

	return nil, enum.ErrNotOpen
}

const GameFilesKey = "{game_files}"

// SetFileList 设置文件列表至redis
func SetFileList(files []*business.GameFile) error {
	return utils.SetCache(GameFilesKey, files, local_cache.NoExpire)
}

func GetFileList() (files []*business.GameFile, err error) {
	err = utils.GetCache(GameFilesKey, &files)
	if err != nil && errors.Is(err, redis.Nil) {
		return []*business.GameFile{}, nil
	}
	return
}

//func InitFileList() error {
//	fileList, err := slotfile.GetAllGameFileList(0, slotfile.SlotRoot)
//	if err != nil {
//		return err
//	}
//	// 截去前缀
//	subLen := len(slotfile.SlotRoot) + 1
//	for _, file := range fileList {
//		if len(file.Path) < subLen {
//			continue
//		}
//		file.Path = file.Path[subLen:]
//	}
//	err = SetFileList(fileList)
//	return err
//}

const CoinTypeKey = "coin_type_key"

func GetCoinMap() (coinMap map[uint][]int64) {
	res, exist := Cache.Get(CoinTypeKey)
	if exist && res != nil {
		return res.(map[uint][]int64)
	}
	var data []*business.Coins
	global.GVA_DB.Select("coin_type", "bet_num").Find(&data)
	coinMap = lo.SliceToMap(data, func(item *business.Coins) (uint, []int64) {
		return item.CoinType, helper.SplitInt[int64](item.BetNum, ",")
	})
	Cache.Set(CoinTypeKey, coinMap, local_cache.NoExpire)
	return
}
