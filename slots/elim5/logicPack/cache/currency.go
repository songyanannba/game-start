package cache

import (
	"elim5/global"
	"elim5/model/business"
	"time"
)

const currencyKey = "currency_info"

func GetCurrencyMap() map[string]*business.Currency {
	var currencyMap map[string]*business.Currency
	info, ok := Cache.Get(currencyKey)
	if !ok {
		// 从数据库获取
		currencyMap = make(map[string]*business.Currency)
		var (
			currencyList []*business.Currency
		)
		global.GVA_READ_DB.Find(&currencyList)
		for _, c := range currencyList {
			currencyMap[c.Name] = c
		}
		//5分钟过期
		Cache.Set(currencyKey, currencyMap, 5*time.Minute)
	} else {
		// 从缓存获取
		currencyMap = info.(map[string]*business.Currency)
	}
	return currencyMap
}

func GetCurrency(name string) *business.Currency {
	return GetCurrencyMap()[name]
}

func GetCurrencyRate(name string) float64 {
	currencyMap := GetCurrencyMap()
	if _, ok := currencyMap[name]; !ok {
		return 1
	} else {
		return currencyMap[name].ExchangeRate
	}
}

func GetCurrencyRateMap() map[string]float64 {
	currencyMap := GetCurrencyMap()
	res := make(map[string]float64)
	for _, c := range currencyMap {
		if c.ExchangeRate == 0 {
			continue
		}
		res[c.Name] = 1 / c.ExchangeRate
	}
	return res
}
