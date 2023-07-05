package base

import (
	"slot6/utils"
	"slot6/utils/helper"
	"strings"
)

type Bet struct {
	Currency string
	Bets     []int32
}

func NewBet(currency string, bets []int32) *Bet {
	return &Bet{
		Currency: currency,
		Bets:     bets,
	}
}

type BetMap struct {
	m map[string]*Bet
}

func NewBetMap(s string) BetMap {
	m := BetMap{m: make(map[string]*Bet)}
	if s == "" {
		return m
	}
	arr := utils.FormatCommand(s)
	for _, betStr := range arr {
		currency, bets, ok := strings.Cut(betStr, ":")
		if !ok {
			continue
		}
		currency = strings.ToUpper(currency)
		m.m[currency] = NewBet(currency, helper.SplitInt[int32](bets, ","))
	}
	return m
}

func (b BetMap) Get(currency string) *Bet {
	currency = strings.ToUpper(currency)
	if bet, ok := b.m[currency]; ok {
		return bet
	}
	if bet, ok := b.m["USD"]; ok {
		return bet
	}
	return NewBet("", []int32{0})
}

func (b BetMap) Check(currency string, bet int64) bool {
	if bet == 0 {
		return false
	}
	return helper.InArr(int32(bet), b.Get(currency).Bets)
}
