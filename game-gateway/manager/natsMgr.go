package manager

import (
	"game-gateway/protoc/pb"
	"github.com/nats-io/nats.go"
	"sync"
)

type natsManager struct {
	Nats       *nats.Conn
	PlayersSub map[string]*nats.Subscription
	Sync       sync.Mutex
}

var NastManager = natsManager{
	Nats:       nil,
	PlayersSub: make(map[string]*nats.Subscription),
	Sync:       sync.Mutex{},
}

func (n *natsManager) Send(where string, msg *pb.NetMessage) {

}
