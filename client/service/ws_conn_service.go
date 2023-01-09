package service

import (
	"client/conf"
	"client/protoc/pb"
	"fmt"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type wsClientService struct {
	sync.Mutex
	context chan []byte
	conn    *websocket.Conn
}

func (ws *wsClientService) Send(msgType int32, userId, serviceId string, message proto.Message) {

	messageMarshal, _ := proto.Marshal(message)
	msg := &pb.NetMessage{
		ServiceId: serviceId,
		UId:       userId,
		Content:   messageMarshal,
		Type:      msgType,
	}

	msgMarshal, _ := proto.Marshal(msg)
	ws.context <- msgMarshal
}

var WsClientService = wsClientService{
	context: make(chan []byte, 1024),
}

func (ws *wsClientService) Start() {
	header := http.Header{}
	header.Add("uid", "syn")
	header.Add("auth", "syn")

	u := url.URL{
		Scheme: "ws",
		Host:   conf.HOST,
		Path:   conf.PATH,
	}
	s := u.String()

	fmt.Println("url str == ", s)

	conn, _, err := websocket.DefaultDialer.Dial(s, nil)
	defer conn.Close()

	ws.conn = conn
	if err != nil {
		fmt.Println("ws dail 服务拨号失败 = ", err)
		return
	}
	//go ws.Read()
	go ws.Write()

	go ws.Test123()

	time.Sleep(100 * time.Second)
}

func (ws *wsClientService) Read() {
	fmt.Println("Read for")

	defer ws.conn.Close()

	for {
		var err error
		mType, msg, err := ws.conn.ReadMessage()
		if mType == websocket.BinaryMessage {
			netMsg := &pb.NetMessage{}
			err = proto.Unmarshal(msg, netMsg)
			if err == nil {
				fmt.Println("wsClientService read = ", netMsg.Type)
				CliHandler.DaYin(netMsg)
			} else {
				fmt.Println("wsClientService read err = ", err)
			}
		}

	}

}

func (ws *wsClientService) Write() {

	for {
		fmt.Println("write for")
		select {
		case context := <-ws.context:
			err := ws.conn.WriteMessage(websocket.BinaryMessage, context)
			if err != nil {
				fmt.Println("ws write err", err)
			} else {
				mm := &pb.NetMessage{}
				proto.Unmarshal(context, mm)
				fmt.Println("ws write succ", mm)
			}
		}
	}

}
func (ws *wsClientService) Test123() {

	request := &pb.GameMessage{
		To:   "123",
		Do:   "234",
		Todo: "4哈哈哈哈6",
	}
	marshal, _ := proto.Marshal(request)

	req := &pb.NetMessage{
		ServiceId: "syn-service",
		UId:       "syn--",
		Content:   marshal,
		Type:      1,
	}
	reqM, _ := proto.Marshal(req)

	ws.context <- reqM
	/*for {
		time.Sleep(30 * time.Second)
	}*/
}
