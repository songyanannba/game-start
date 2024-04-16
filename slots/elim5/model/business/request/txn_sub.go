package request

import (
	"elim5/model/business"
	"elim5/model/common/request"
	"time"
)

type TxnSubSearch struct {
	business.TxnSub
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartBet       *float64   `json:"startBet" form:"startBet"`
	EndBet         *float64   `json:"endBet" form:"endBet"`
	StartRaise     *float64   `json:"startRaise" form:"startRaise"`
	EndRaise       *float64   `json:"endRaise" form:"endRaise"`
	StartWin       *float64   `json:"startWin" form:"startWin"`
	EndWin         *float64   `json:"endWin" form:"endWin"`
	PlayerName     string     `json:"playerName" form:"playerName"`
	UserId         uint       `json:"userId" form:"userId"`
	GameId         int        `json:"gameId" form:"gameId"`
	request.PageInfo
}
