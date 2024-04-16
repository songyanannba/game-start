package base

import (
	"elim5/pbs/game"
)

func (t *Tag) ToBaseCard() *game.BaseCard {
	return &game.BaseCard{
		Id:     int32(t.Id),
		X:      int32(t.X),
		Y:      int32(t.Y),
		Mul:    int32(t.Multiple),
		IsWild: t.IsWild,
		IsPay:  t.IsPayTable,
	}
}
