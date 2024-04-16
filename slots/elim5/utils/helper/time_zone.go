package helper

import (
	"elim5/global"
	"fmt"
	"go.uber.org/zap"
	"time"
)

// TimeToUTC 时区强制转换为UTC时区 2021-01-01 00:00:00 +0800 CST -> 2021-01-01 00:00:00 +0000 UTC
func TimeToUTC(t time.Time) time.Time {
	_, offset := t.Zone()
	return t.Add(time.Duration(offset) * time.Second).In(time.UTC)
}

// TimeOffset 时区强制转换为目标时区
func TimeOffset(date time.Time, timeZone string, flip bool) time.Time {
	offset := GetOffsetInt(timeZone)
	if flip {
		offset = -offset
	}
	return date.Add(time.Duration(offset) * time.Second)
}

// GetOffsetInt 计算目标时区与服务器本地时区的偏移量
func GetOffsetInt(timeZone string) int {
	//global.GVA_LOG.Info("timeZone", zap.String("timeZone", timeZone))
	targetLocation, err := time.LoadLocation(timeZone)
	if err != nil {
		global.GVA_LOG.Error("无法加载目标时区:", zap.Error(err))
		return 0
	}

	now := time.Now()
	targetTime := now.In(targetLocation)

	// 计算目标时区与服务器本地时区的偏移量
	_, offset := targetTime.Zone()
	_, nowOffset := now.Zone()

	return nowOffset - offset
}

// GetOffsetString 将时区偏移量格式化为"-07:00"字符串
func GetOffsetString(timeZone string) string {
	offset := GetOffsetInt(timeZone)
	return fmt.Sprintf("%s%02d:%02d", getOffsetSign(offset), int32(Abs(float64(offset/3600))), Abs((offset%3600)/60))
}

// 根据偏移量的正负返回符号
func getOffsetSign(offset int) string {
	if offset > 0 {
		return "-"
	}
	return "+"
}
