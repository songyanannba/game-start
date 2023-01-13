package service

import (
	"fmt"
	"game-player/conf"
	"game-player/protoc/pb"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"

	"sync"
	"time"
)

var NatsManager = natsManager{
	Nats:          nil,
	PlayerSubBack: nil,
	PlayersSubMap: make(map[string]*nats.Subscription),
}

type natsManager struct {
	Nats          *nats.Conn
	PlayerSubBack *nats.Subscription
	PlayersSubMap map[string]*nats.Subscription
	Sync          sync.Mutex
}

func (n *natsManager) Start() {

	connect, err := nats.Connect(
		fmt.Sprintf("nats://127.0.0.1:%d", 4222),
		nats.MaxReconnects(10),
		nats.RetryOnFailedConnect(true),
		nats.ReconnectWait(15*time.Millisecond),
		nats.DisconnectErrHandler(func(_ *nats.Conn, _ error) {

		}),
	)
	if err != nil {
		fmt.Println("nats conn err = ", err)
		return
	}
	//defer connect.Close()

	n.Nats = connect

	n.SubBack(conf.PlayerConf.ServiceId)

}

func (n *natsManager) SubBack(serviceID string) {

	n.PlayerSubBack, _ = n.Nats.Subscribe(serviceID , func(msg *nats.Msg) {
		netMsg := &pb.NetMessage{}
		err := proto.Unmarshal(msg.Data, netMsg)
		if err != nil {
			fmt.Println("SubBack err ", err )
		}
		fmt.Println("netMessage == " , netMsg)

		//服务调用

	})

}
