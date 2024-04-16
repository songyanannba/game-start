package cache

import (
	"context"
	"elim5/global"
	"elim5/pbs/game"
	"elim5/utils/helper"
	"errors"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"strconv"
	"time"
)

const GameProcessKey = "{game_process}"

func GetGameProcessKey(userId int64) string {
	return strconv.Itoa(int(userId))
}

// GetGameProcess 获取游戏进程
func GetGameProcess(userId int64) (*game.GameProcess, error) {
	process := &game.GameProcess{}
	res, err := global.GVA_REDIS.HGet(context.Background(), GameProcessKey, GetGameProcessKey(userId)).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return process, nil
		}
		global.GVA_LOG.Error("获取游戏进程失败", zap.Error(err))
		return process, err
	}

	err = proto.Unmarshal(res, process)
	if err != nil {
		global.GVA_LOG.Error("反序列化游戏进程失败", zap.Error(err))
		return process, err
	}
	return process, nil
}

func GetAllGameProcess() (map[uint]*game.GameProcess, error) {
	res, err := global.GVA_REDIS.HGetAll(context.Background(), GameProcessKey).Result()
	if err != nil {
		if err != nil {
			global.GVA_LOG.Error("获取游戏进程失败", zap.Error(err))
		}
		return nil, err
	}
	m := map[uint]*game.GameProcess{}
	for uid, v := range res {
		process := &game.GameProcess{}
		err = proto.Unmarshal([]byte(v), process)
		if err != nil {
			global.GVA_LOG.Error("解析游戏进程proto失败", zap.Error(err))
			continue
		}

		m[uint(helper.Atoi(uid))] = process
	}
	return m, nil
}

// SetGameProcess 设置游戏进程
func SetGameProcess(userId int64, process *game.GameProcess) error {
	key := GetGameProcessKey(userId)
	process.CreateTime = time.Now().Unix()
	data, err := proto.Marshal(process)
	if err != nil {
		return err
	}

	err = global.GVA_REDIS.HMSet(context.Background(), GameProcessKey, key, data).Err()
	return nil
}

// DelGameProcess 删除游戏进程
func DelGameProcess(userId int64) (err error) {
	return global.GVA_REDIS.HDel(context.Background(), GameProcessKey, GetGameProcessKey(userId)).Err()
}
