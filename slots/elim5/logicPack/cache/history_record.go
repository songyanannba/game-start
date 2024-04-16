package cache

import (
	"context"
	"elim5/global"
	"elim5/model/business"
	"elim5/pbs/common"
	"elim5/utils"
	"elim5/utils/helper"
	"fmt"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// 以下为菜单缓存

func GetRecordMenuKey(userId uint, req *common.RecordMenuReq) string {
	return "{record_menu}:" + req.TimeZone + strconv.Itoa(int(userId)) + ":" + strconv.Itoa(int(req.GameId)) + ":" + req.Date
}

// GetRecordMenuCache  获取用户机台游戏记录缓存
func GetRecordMenuCache(userId uint, req *common.RecordMenuReq) (recordMenu *RecordMenu, err error) {
	key := GetRecordMenuKey(userId, req)
	recordMenu = &RecordMenu{}
	var result []byte
	result, err = global.GVA_REDIS.Get(context.Background(), key).Bytes()
	if err != nil {
		return
	}
	err = global.Json.Unmarshal(result, &recordMenu)
	if err != nil {
		global.GVA_LOG.Error("GetRecordMenuCache err", zap.Error(err))
		return
	}
	return
}

// SetRecordMenuCache  设置用户机台游戏记录缓存
func SetRecordMenuCache(userId uint, req *common.RecordMenuReq, recordMenu *RecordMenu) (err error) {
	key := GetRecordMenuKey(userId, req)
	var data []byte
	for i, int32s := range recordMenu.DayHourMap {
		recordMenu.DayHourMap[i] = helper.Distinct(int32s)
	}
	data, err = global.Json.Marshal(recordMenu)
	err = global.GVA_REDIS.Set(context.Background(), key, data, 1*time.Hour).Err()
	return
}

type RecordMenu struct {
	DayHourMap map[int32][]int32 `json:"day_hour_map"`
	MaxTime    time.Time         `json:"max_time"`
}

func GetRecordMenu(userId uint, req *common.RecordMenuReq) (dayHourMap map[int32][]int32, err error) {
	var (
		timeZone   *time.Location
		recordMenu *RecordMenu
	)
	// 加载目标时区
	timeZone, err = time.LoadLocation(req.TimeZone)
	if err != nil {
		return
	}

	//获取开始时间 2020-01
	var startTime time.Time
	startTime, err = time.ParseInLocation("2006-01", req.Date, timeZone)
	if err != nil {
		return
	}

	// 从缓存获取菜单
	recordMenu, err = GetRecordMenuCache(userId, req)
	dayHourMap = recordMenu.DayHourMap

	if err != nil || recordMenu.MaxTime.Before(startTime) || len(recordMenu.DayHourMap) == 0 {
		//无缓存查数据库
		recordMenu, err = DbGetRecordMenu(userId, req, startTime, startTime, timeZone)
		if err != nil {
			return
		}
		//更新缓存
		err = SetRecordMenuCache(userId, req, recordMenu)
		if err != nil {
			global.GVA_LOG.Error("SetRecordMenuCache err", zap.Error(err))
		}
		dayHourMap = recordMenu.DayHourMap
		return
	}

	//如果不是当前月份,且有缓存返回缓存
	if recordMenu.MaxTime.UTC().Month() != time.Now().UTC().Month() {
		return
	}

	//如果是当前月份,且最后一小时有数据,返回缓存

	if recordMenu.MaxTime.UTC().Year() == time.Now().UTC().Year() &&
		recordMenu.MaxTime.UTC().Month() == time.Now().UTC().Month() &&
		recordMenu.MaxTime.UTC().Day() == time.Now().UTC().Day() &&
		recordMenu.MaxTime.UTC().Hour() == time.Now().UTC().Hour() {
		return
	}

	//如果是当前月份,且最后一小时没有数据,更新缓存
	addRecordMenu := &RecordMenu{}
	addRecordMenu, err = DbGetRecordMenu(userId, req, startTime, recordMenu.MaxTime, timeZone)
	if err != nil {
		return
	}
	for k, v := range addRecordMenu.DayHourMap {
		for _, i2 := range v {
			recordMenu.DayHourMap[k] = helper.Distinct(append(recordMenu.DayHourMap[k], i2))
		}
	}
	recordMenu.MaxTime = addRecordMenu.MaxTime
	err = SetRecordMenuCache(userId, req, recordMenu)
	if err != nil {
		global.GVA_LOG.Error("SetRecordMenuCache err", zap.Error(err))
	}
	return
}

func DbGetRecordMenu(userId uint, req *common.RecordMenuReq, startTime, minTime time.Time, timeZone *time.Location) (recordMenu *RecordMenu, err error) {
	var (
		endTime time.Time
		times   []time.Time
	)
	recordMenu = &RecordMenu{
		DayHourMap: make(map[int32][]int32),
	}
	//获取最大时间
	endTime = startTime.AddDate(0, 1, 0).Add(-time.Nanosecond)

	err = global.GVA_READ_DB.Model(&business.SlotRecord{}).
		Select("created_at").
		Where("created_at BETWEEN ? AND ?", minTime.UTC(), endTime.UTC()).
		Where("user_id = ? ", userId).
		Where("slot_id = ? ", req.GameId).
		Scan(&times).Error
	if err != nil {
		return
	}

	maxTime := lo.MaxBy(times, func(a, b time.Time) bool {
		return a.Before(b)
	})
	recordMenu.MaxTime = maxTime.UTC()
	days := lo.GroupBy(times, func(item time.Time) int {
		return item.Day()
	})
	for day, hours := range days {
		recordMenu.DayHourMap[int32(day)] = helper.Distinct(lo.FilterMap[time.Time, int32](hours, func(item time.Time, index int) (int32, bool) {
			return int32(item.UTC().In(timeZone).Hour()), true
		}))
	}
	return
}

// 以下为列表缓存

func GetRecordListKey(userId uint, gameId int32, date string) string {
	return fmt.Sprintf("{record_list}:%d:%d:%s", userId, gameId, date)
}

func GetRecordListCache(userId uint, gameId int32, date string) (ack *common.RecordListAck, err error) {
	key := GetRecordListKey(userId, gameId, date)
	ack = &common.RecordListAck{}
	err = utils.GetCache(key, &ack)
	if err != nil {
		return nil, err
	}
	return
}

func SetRecordListCache(userId uint, gameId int32, date string, ack *common.RecordListAck) (err error) {
	key := GetRecordListKey(userId, gameId, date)
	randSec := helper.RandInt(120)
	randSec += 60
	err = utils.SetCache(key, ack, time.Duration(randSec)*time.Second)
	if err != nil {
		return err
	}
	return
}

func DeleteRecordListCache(userId uint, gameId int32, date string) (err error) {
	err = utils.DelCache(GetRecordListKey(userId, gameId, date))
	if err != nil {
		return err
	}
	return
}
