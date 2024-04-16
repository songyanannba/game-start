package cache

import (
	"context"
	"elim5/enum"
	"elim5/global"
	"strconv"
	"time"
)

type GambleStat struct {
	State  int `json:"state"` //0  未完成 1 完成
	Gamble int `json:"gamble"`
	SlotId int `json:"slot_id"`
}

func GetGambleStatsKey(userId int64) string {
	return "{user_gamble_status}:" + strconv.Itoa(int(userId))
}

func GetGambleTtl(userId int64) (time.Duration, error) {
	key := GetGambleStatsKey(userId)
	result, err := global.GVA_REDIS.TTL(context.Background(), key).Result()
	return result, err
}

func GetGambleStats(userId int64) (gss []*GambleStat, err error) {
	key := GetGambleStatsKey(userId)
	var result []byte
	result, err = global.GVA_REDIS.Get(context.Background(), key).Bytes()
	if err != nil {
		return
	}
	err = global.Json.Unmarshal(result, &gss)
	return
}

func SetGambleStats(userId int64, gs *GambleStat) (err error) {
	key := GetGambleStatsKey(userId)
	gambleStats, err := GetGambleStats(userId)
	if err == nil && len(gambleStats) > 0 {
		stat := gambleStats[len(gambleStats)-1]
		if gs.Gamble > stat.Gamble {
			gambleStats = append(gambleStats, gs)
		} else if gs.State == 1 && gs.Gamble == stat.Gamble {
			stat.State = 1
		}
	}
	if err != nil {
		if err.Error() == enum.RedisNil {
			gambleStats = append(gambleStats, gs)
		}
	}

	data, err := global.Json.Marshal(gambleStats)
	if err != nil {
		return
	}
	err = global.GVA_REDIS.Set(context.Background(), key, data, time.Hour*24*15).Err()
	return
}

func DelGambleStats(userId int64) (err error) {
	return global.GVA_REDIS.Del(context.Background(), GetGambleStatsKey(userId)).Err()
}

//单个

//func GetUserGambleKey(userId int64) string {
//	return "{user_gamble}:" + strconv.Itoa(int(userId))
//}
//
//func GetUserGamble(userId int64) (ack *GambleStat, err error) {
//	ack = &GambleStat{}
//	key := GetUserGambleKey(userId)
//	var result []byte
//	result, err = global.GVA_REDIS.Get(context.Background(), key).Bytes()
//	if err != nil {
//		return
//	}
//	err = global.Json.Unmarshal(result, ack)
//	return
//}
//
//func SetUserGamble(userId int64, ack *GambleStat) (err error) {
//	key := GetUserGambleKey(userId)
//	var data []byte
//	data, err = global.Json.Marshal(ack)
//	if err != nil {
//		return
//	}
//	err = global.GVA_REDIS.Set(context.Background(), key, data, time.Hour*24*15).Err()
//	return
//}
//
//func DelUserGamble(userId int64) (err error) {
//	return global.GVA_REDIS.Del(context.Background(), GetUserGambleKey(userId)).Err()
//}

//如果普通砖存在神秘标签 记录历史位置

type MysteryTagCol struct {
	UserId         int
	SlotId         int //机器编号
	Bet            int
	Type           int
	Trigger        int    //类型 1:普通砖;2:free_spin
	ProgressStatus int8   //状态 1 未完成 2 完成 （长标签的进度）
	Layout         string //排布
	Dir            int    //0 下 1 上
	Step           int
}

func GetUserMysteryTagColKey(userId, bet int64) string {
	return "{user_mystery_col_bet}:" + strconv.Itoa(int(userId)) + strconv.Itoa(int(bet))
}

func GetMysteryTagCol(userId, bet int64) (ssn *MysteryTagCol, err error) {
	ssn = &MysteryTagCol{}
	key := GetUserMysteryTagColKey(userId, bet)
	var result []byte
	result, err = global.GVA_REDIS.Get(context.Background(), key).Bytes()
	if err != nil {
		return
	}
	err = global.Json.Unmarshal(result, ssn)
	return
}

func SetMysteryTagCol(userId int64, ack *MysteryTagCol) (err error) {
	key := GetUserMysteryTagColKey(userId, int64(ack.Bet))
	var data []byte
	data, err = global.Json.Marshal(ack)
	if err != nil {
		return
	}
	err = global.GVA_REDIS.Set(context.Background(), key, data, time.Hour*24*15).Err()
	return
}

func DelMysteryTagCol(userId, bet int64) (err error) {
	return global.GVA_REDIS.Del(context.Background(), GetUserMysteryTagColKey(userId, bet)).Err()
}
