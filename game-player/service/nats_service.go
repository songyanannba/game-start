package service

import (
	"fmt"
	"github.com/nats-io/nats.go"

	"sync"
	"time"
)

var NatsManager = natsManager{
	Nats:       nil,
	PlayersSub: make(map[string]*nats.Subscription),
}

type natsManager struct {
	Nats       *nats.Conn
	PlayersSub map[string]*nats.Subscription
	Sync       sync.Mutex
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
	defer connect.Close()

	n.Nats = connect

}
