package service

import (
	"game-player/db"
	"game-player/protoc/pb"
	"github.com/nats-io/nats.go"
	"sync"
	"time"
)

type playerSpace struct {
	PlayerInfo  *db.PlayerInfo
	CurrentTime time.Time
}

var PlayerService = playerService{
	Players:       make(map[string]*playerSpace),
	PlayerHandler: make(map[int32]*func(ps *playerSpace, msg *pb.NetMessage)),
}

type playerService struct {
	Players       map[string]*playerSpace
	PlayerHandler map[int32]*func(ps *playerSpace, msg *pb.NetMessage)

	sync sync.Mutex
}

func (ps *playerService) Start() {

	NatsManager.Nats.Subscribe("", func(msg *nats.Msg) {

	})
	//注册handler
}
