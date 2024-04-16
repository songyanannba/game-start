package cache

import (
	"context"
	"elim5/enum"
	"elim5/global"
	"elim5/model/business"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"strconv"
)

const GameRecord = "{game_record}"

func GetGameRecordKey(recordId uint) string {
	return strconv.Itoa(int(recordId))
}

// GetGameRecord 获取游戏记录
func GetGameRecord(recordId uint) (record *business.SlotRecord, err error) {
	var res []byte
	record = &business.SlotRecord{}
	res, err = global.GVA_REDIS.HGet(context.Background(), GameRecord, GetGameRecordKey(recordId)).Bytes()
	if err == nil {
		err = global.Json.Unmarshal(res, record)
		if err != nil {
			global.GVA_LOG.Error("反序列化游戏记录失败", zap.Error(err))
		}
		return
	}

	err = global.GVA_DB.First(record, "id = ?", recordId).Error
	if err != nil {
		global.GVA_LOG.Error("获取游戏记录失败", zap.Error(err))
		return nil, enum.ErrCommon
	}
	err = SetGameRecord(record)
	return
}

// FallAllGameRecord 落盘所有特殊流程的游戏记录
func FallAllGameRecord() error {
	res, err := global.GVA_REDIS.HGetAll(context.Background(), GameRecord).Result()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			return nil
		}
		global.GVA_LOG.Error("获取游戏记录失败", zap.Error(err))
		return err
	}

	for _, v := range res {
		var record *business.SlotRecord
		record, err = ParseGameRecord(v)
		if err != nil {
			return err
		}

		err = global.GVA_DB.Save(record).Error
		if err != nil {
			global.GVA_LOG.Error("保存游戏记录失败", zap.Error(err))
			return err
		}

		err = DelGameRecord(record.ID)
		if err != nil {
			global.GVA_LOG.Error("删除游戏记录失败", zap.Error(err))
			return err
		}
	}
	return nil
}

func ParseGameRecord(s string) (*business.SlotRecord, error) {
	var record *business.SlotRecord
	err := global.Json.Unmarshal([]byte(s), &record)
	if err != nil {
		global.GVA_LOG.Error("反序列化游戏记录失败", zap.Error(err))
		return nil, err
	}
	return record, nil
}

func SetGameRecord(record *business.SlotRecord) error {
	s, err := global.Json.MarshalToString(record)
	if err != nil {
		global.GVA_LOG.Error("序列化游戏记录失败", zap.Error(err))
		return err
	}
	err = global.GVA_REDIS.HSet(context.Background(), GameRecord, GetGameRecordKey(record.ID), s).Err()
	if err != nil {
		global.GVA_LOG.Error("缓存游戏记录失败", zap.Error(err))
	}
	return err
}

func DelGameRecord(recordId ...uint) error {
	fields := make([]string, len(recordId))
	for i, v := range recordId {
		fields[i] = GetGameRecordKey(v)
	}
	return global.GVA_REDIS.HDel(context.Background(), GameRecord, fields...).Err()
}
