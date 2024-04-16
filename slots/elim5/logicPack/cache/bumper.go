package cache

import (
	"context"
	"elim5/global"
	"elim5/model/business"
	"elim5/utils/conver"
	"elim5/utils/helper"
	"strconv"
	"strings"
	"time"
)

const BumperKey = "{bumper}"

func GetBumperKey(slotId uint, merchantId uint, currency string) []string {
	id := helper.Itoa(slotId) + "_" + helper.Itoa(merchantId)
	return []string{id + "_" + currency + "_b", id + "_" + currency + "_w"}
}

// GetBumper 获取货币当前的保险杠值 [总押注、总赢取] merchantId为0时获取全局
func GetBumper(slotId uint, merchantId uint, currency string) (int64, int64, error) {
	keys := GetBumperKey(slotId, merchantId, currency)
	arr, err := global.GVA_REDIS.HMGet(context.Background(), BumperKey, keys...).Result()
	if err != nil {
		return 0, 0, err
	}

	// map中没有数据时从数据库中获取
	if arr[0] == nil && arr[1] == nil {
		// 如果商户id不为0说明商户的保险杠数据没有初始化
		if merchantId != 0 {
			// 初始化商户数据
			var data []*bumperData
			data, err = InitBumperData(merchantId)
			if err != nil {
				return 0, 0, err
			}
			// 初始化成功时返回对应的数据
			for _, v := range data {
				if v.SlotId == slotId && v.Currency == currency {
					return v.Bets, v.Wins, nil
				}
			}
		}
		return 0, 0, nil
	}

	return conver.Int64Must(arr[0]), conver.Int64Must(arr[1]), nil
}

type BumperData struct {
	SlotId     uint
	MerchantId uint
	Currency   string
	Bets       int64
	Wins       int64
}

// GetAllBumper 获取所有货币的保险杠值
func GetAllBumper() (map[string]*BumperData, error) {
	m, err := global.GVA_REDIS.HGetAll(context.Background(), BumperKey).Result()
	if err != nil {
		return nil, err
	}
	dataMap := map[string]*BumperData{}
	for key, val := range m {
		arr := strings.Split(key, "_")
		if len(arr) != 4 {
			continue
		}
		currency := arr[2]

		mKey := arr[0] + "_" + arr[1] + "_" + currency

		data := dataMap[mKey]
		if data == nil {
			data = &BumperData{
				SlotId:     uint(helper.Atoi(arr[0])),
				MerchantId: uint(helper.Atoi(arr[1])),
				Currency:   currency,
			}
			dataMap[mKey] = data
		}
		if arr[3] == "b" {
			data.Bets, _ = strconv.ParseInt(val, 10, 64)
		} else if arr[3] == "w" {
			data.Wins, _ = strconv.ParseInt(val, 10, 64)
		}
	}
	return dataMap, nil
}

// IncBumper 增加保险杠数据 当商户保险杠配置存在时增加商户保险杠数据
func IncBumper(slotId, merchantId uint, currency string, bet int64, win int64) (int64, int64, error) {
	data := NewIncBumperData(slotId, 0, currency, bet, win)
	if merchantId > 0 && MerchantInBumperConfig(merchantId) {
		data = append(NewIncBumperData(slotId, merchantId, currency, bet, win), data...)
	}
	arr, err := AtomicOperate(data...).Int64Slice()
	if err != nil {
		return 0, 0, err
	}
	return arr[0], arr[1], nil
}

// NewIncBumperData 生成保险杠增量数据
func NewIncBumperData(slotId, merchantId uint, currency string, bet int64, win int64) []*Eval {
	keys := GetBumperKey(slotId, merchantId, currency)
	return []*Eval{
		{BumperKey, "HIncrBy", []any{keys[0], bet}},
		{BumperKey, "HIncrBy", []any{keys[1], win}},
	}
}

func SetBumper(slotId, merchantId uint, currency string, bet int64, win int64) error {
	keys := GetBumperKey(slotId, merchantId, currency)
	err := global.GVA_REDIS.HSet(context.Background(), BumperKey, keys[0], bet, keys[1], win).Err()
	return err
}

func ClearBumper() error {
	err := global.GVA_REDIS.Del(context.Background(), BumperKey).Err()
	return err
}

type bumperData struct {
	SlotId   uint   `gorm:"column:slot_id"`
	Currency string `gorm:"column:currency"`
	Bets     int64  `gorm:"column:bets"`
	Wins     int64  `gorm:"column:wins"`
}

// InitBumperData 通过商户初始化全局的保险杠数据
func InitBumperData(merchantId uint) ([]*bumperData, error) {
	if merchantId == 0 {
		global.GVA_REDIS.Del(context.Background(), BumperKey)
	} else {
		lock := "bumperDataLock:" + helper.Itoa(merchantId)
		if err := global.GVA_REDIS.SetNX(context.Background(), lock, 1, time.Second*15).Err(); err != nil {
			global.GVA_LOG.Warn("并发获取保险杠...")
			return nil, err
		}
	}
	now := helper.GetDateByMonth(time.Now())
	var data []*bumperData
	q := global.GVA_READ_DB.Table(business.SlotRecord{}.TableName()).
		Select("`slot_id`, `currency`, SUM(`total_bet`) as bets, SUM(`gain`) as wins").
		Where("created_at between ? and ?", now, now.AddDate(0, 1, 0).Add(-time.Second))
	if merchantId != 0 {
		q = q.Where("merchant_id = ?", merchantId)
	}
	err := q.Group("slot_id, currency").Find(&data).Error
	if err != nil {
		global.GVA_LOG.Error("init bumper data error: " + err.Error())
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	var values []any
	for _, v := range data {
		keys := GetBumperKey(v.SlotId, merchantId, v.Currency)
		values = append(values, keys[0], v.Bets, keys[1], v.Wins)
	}
	// 一次性将所有数据写入redis
	err = global.GVA_REDIS.HSet(context.Background(), BumperKey, values...).Err()
	if err != nil {
		global.GVA_LOG.Error("init bumper data error: " + err.Error())
	}
	return data, err
}
