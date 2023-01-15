package manager

import (
	"fmt"
	"game-gateway/protoc/pb"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"sync"
	"time"
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

/*func (n *natsManager) SubTopic(topic string) {

	 n.Nats.Subscribe(topic, func(msg *nats.Msg) {
		netMsg := &pb.NetMessage{}
		err := proto.Unmarshal(msg.Data, netMsg)
		if err != nil {
			fmt.Println("SubTopic err ", err)
		}
		fmt.Println("接收到 服务端 发来的信息 netMessage == ", netMsg)

		// 找到userid
		//发 往 conn manager sendOut
		//然后 长链接 管理器 找到 user对应的 连接
		service.GameClientManager.SendOutBack(netMsg)

	})
}*/
