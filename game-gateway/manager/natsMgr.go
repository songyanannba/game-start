package manager

import (
	"fmt"
	"game-gateway/protoc/pb"
	"github.com/golang/protobuf/proto"
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

func (n *natsManager) servicePub(where string, msg *pb.NetMessage) {

	marshal, _ := proto.Marshal(msg)

	err := n.Nats.Publish(where, marshal)

	if err != nil {
		fmt.Println("natsManager err : ", err)
	}
}

func (n *natsManager) serviceSub(where string) {

	_, err := n.Nats.Subscribe(where, func(msg *nats.Msg) {

	})
	if err != nil {
		fmt.Println("serviceSub err = ", err)
		return
	}

}

func (n *natsManager) Send(where string, msg *pb.NetMessage) {
	n.servicePub(where, msg)
}

func (n *natsManager) SubTopic(topic string, f func(msg *pb.NetMessage)) {

}
