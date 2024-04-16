package component

import (
	"elim5/pbs/game"
)

func (s *Spin) GetBaseSpinStep() *game.BaseSpinStep {
	baseSpinStep := &game.BaseSpinStep{
		Type:  int32(s.Type()),
		Id:    int32(s.Id),
		Pid:   int32(s.ParentId),
		Which: int32(s.Which),
		//JackpotId: int32(s.Jackpot.Id),
		Gain: int64(s.Gain),
	}

	if s.Jackpot != nil {
		baseSpinStep.JackpotId = int32(s.Jackpot.Id)
	}
	return baseSpinStep
}
