package service

import (
	"client2/protoc/pb"
	"fmt"
)

type cliHandler struct {
}

var CliHandler = &cliHandler{}

func (ch *cliHandler) Start() {

}

func (ch *cliHandler) DaYin(msg *pb.NetMessage) {

	fmt.Println("dayin msg == ", msg)

}
