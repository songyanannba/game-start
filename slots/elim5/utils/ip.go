package utils

import (
	"context"
	"elim5/global"
	"go.uber.org/zap"
	"net"
	"strconv"
	"time"
)

type Ip struct {
	Ip          string  `json:"ip"`
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Message     string  `json:"message"`
	Country     string  `json:"country"`     // 国家/地区名称 China
	CountryCode string  `json:"countryCode"` // 两个字母的国家代码 CN
	Region      string  `json:"region"`      // 地区/州短码 HB
	RegionName  string  `json:"regionName"`  // 地区/州 Hubei
	City        string  `json:"city"`        // 城市 Wuhan
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"` // 时区 Asia/Shanghai
	Isp         string  `json:"isp"`      // ISP name
	Org         string  `json:"org"`      // Organization name
	As          string  `json:"as"`
}

func GetIpKey(ip string) string {
	return "{ip_info}:" + ip
}

func GetCountryCodeByIp(ip string) (string, error) {
	key := GetIpKey(ip)
	country, err := global.GVA_REDIS.Get(context.Background(), key).Result()
	if err == nil {
		return country, nil
	}

	var res []byte
	url := "https://pro.ip-api.com/json/" + ip + "?key=" + global.GVA_CONFIG.Keys.IpApiKey
	res, err = NewGurl("GET", url).Set(Option{SkipVerify: true}).Do()
	if err != nil {
		global.GVA_LOG.Error("GetCountryByIp", zap.Error(err))
		return "", err
	}

	var ipInfo Ip
	err = global.Json.Unmarshal(res, &ipInfo)
	if err != nil {
		global.GVA_LOG.Error("Unmarshal GetCountryByIp", zap.Error(err))
		return "", err
	}

	if ipInfo.Status != "success" {
		global.GVA_LOG.Warn("GetCountryByIp", zap.String("ip", ip), zap.String("status", ipInfo.Status), zap.String("msg", ipInfo.Message))
	}

	_ = global.GVA_REDIS.Set(context.Background(), key, ipInfo.CountryCode, 20*time.Minute).Err()
	return ipInfo.CountryCode, nil
}

// SpliteAddress 将普通地址格式(host:port)拆分
func SpliteAddress(addr string) (host string, port int, err error) {
	var portStr string
	host, portStr, err = net.SplitHostPort(addr)
	if err != nil {
		return "", 0, err
	}
	port, err = strconv.Atoi(portStr)
	if err != nil {
		return "", 0, err
	}
	return
}
