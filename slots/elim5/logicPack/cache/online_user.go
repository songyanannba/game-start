package cache

import (
	"context"
	"elim5/global"
	"elim5/utils/helper"
	"go.uber.org/zap"
	"time"
)

const RedisOnlineUserKey = "{online_user}"

// CountOnlineUser 统计在线用户
func CountOnlineUser() (int64, error) {
	count, err := global.GVA_REDIS.ZCard(context.Background(), RedisOnlineUserKey).Result()
	if err != nil {
		global.GVA_LOG.Error("countOnlineUser", zap.Error(err))
		return 0, err
	}
	return count, nil
}

// GetOnlineUser 获取在线用户
func GetOnlineUser(offset ...int64) ([]uint, error) {
	var (
		start int64
		stop  int64 = -1
	)
	if len(offset) == 2 {
		start = offset[0]
		stop = start + offset[1] - 1
	}
	var userIds []uint
	err := global.GVA_REDIS.ZRange(context.Background(), RedisOnlineUserKey, start, stop).ScanSlice(&userIds)
	if err != nil {
		global.GVA_LOG.Error("getOnlineUser", zap.Error(err))
		return nil, err
	}
	return userIds, nil
}

// SetOnlineUser 设置在线用户
//func SetOnlineUser(userIds ...int64) error {
//	if len(userIds) == 0 {
//		return nil
//	}
//
//	expireTime := float64(time.Now().Add(time.Second * 90).Unix())
//
//	var members []redis.Z
//	for _, userId := range userIds {
//		members = append(members, redis.Z{
//			Score:  expireTime,
//			Member: userId,
//		})
//	}
//
//	err := global.GVA_REDIS.ZAdd(context.Background(), RedisOnlineUserKey, members...).Err()
//	if err != nil {
//		global.GVA_LOG.Error("setOnlineUser", zap.Error(err))
//		return err
//	}
//	return nil
//}

// DelOnlineUser 删除在线用户
func DelOnlineUser(userIds ...int64) error {
	if len(userIds) == 0 {
		return nil
	}

	var members []interface{}
	for _, userId := range userIds {
		members = append(members, userId)
	}

	err := global.GVA_REDIS.ZRem(context.Background(), RedisOnlineUserKey, members...).Err()
	if err != nil {
		global.GVA_LOG.Error("delOnlineUser", zap.Error(err))
		return err
	}
	return nil
}

// ClearExpireOnlineUser  清除过期的在线用户
func ClearExpireOnlineUser() error {
	err := global.GVA_REDIS.ZRemRangeByScore(context.Background(), RedisOnlineUserKey, "0", helper.Itoa(time.Now().Unix())).Err()
	if err != nil {
		global.GVA_LOG.Error("clearOnlineUser", zap.Error(err))
		return err
	}
	return nil
}
