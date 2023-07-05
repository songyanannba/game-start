package component

import "github.com/shopspring/decimal"

func (s *Spin) jackpotSum(jackpot *JackpotData) float64 {
	// jackpot值改为固定 只需取最终值即可
	s.Gain = int(decimal.NewFromInt(int64(s.Bet)).Mul(decimal.NewFromFloat(jackpot.End)).IntPart())
	return jackpot.End
}
