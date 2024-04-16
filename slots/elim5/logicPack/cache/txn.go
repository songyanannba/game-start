package cache

import (
	"context"
	"elim5/enum"
	"elim5/global"
	"elim5/model/business"
	"errors"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"strconv"
)

const TxnKey = "{txn}"

func GetUserTxnKey(userID uint) string {
	return strconv.Itoa(int(userID))
}

func GetTxn(userID uint) (*business.Txn, error) {
	res, err := global.GVA_REDIS.HGet(context.Background(), TxnKey, GetUserTxnKey(userID)).Bytes()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			global.GVA_LOG.Error("获取注单失败", zap.Error(err))
		}
		return nil, enum.ErrSlotError
	}
	var txn business.Txn
	err = global.Json.Unmarshal(res, &txn)
	if err != nil {
		global.GVA_LOG.Error("反序列化注单失败", zap.Error(err))
		return nil, err
	}
	return &txn, nil
}

func GetAllTxn() ([]*business.Txn, error) {
	res, err := global.GVA_REDIS.HGetAll(context.Background(), TxnKey).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			global.GVA_LOG.Error("获取注单失败", zap.Error(err))
		}
		return nil, err
	}
	var txnList []*business.Txn
	for _, v := range res {
		var txn *business.Txn
		err = global.Json.UnmarshalFromString(v, &txn)
		if err != nil {
			global.GVA_LOG.Error("反序列化注单失败", zap.Error(err))
			continue
		}
		txnList = append(txnList, txn)
	}
	return txnList, nil
}

func SetTxn(txn *business.Txn) error {
	s, err := global.Json.MarshalToString(txn)
	if err != nil {
		global.GVA_LOG.Error("序列化注单失败", zap.Error(err))
		return err
	}
	err = global.GVA_REDIS.HSet(context.Background(), TxnKey, GetUserTxnKey(txn.UserId), s).Err()
	if err != nil {
		global.GVA_LOG.Error("缓存注单失败", zap.Error(err))
	}
	return err
}

func DelTxn(userId uint) error {
	return global.GVA_REDIS.HDel(context.Background(), TxnKey, GetUserTxnKey(userId)).Err()
}
