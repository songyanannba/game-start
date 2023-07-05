package service

import (
	"client/protoc/pb"
	"sync"
)

type commonService struct {
	sync.Mutex
	HandlerMap map[int32]func(msg *pb.NetMessage)
}

var CommonService = commonService{
	HandlerMap: make(map[int32]func(msg *pb.NetMessage)),
}

func (cs *commonService) First() {

}

func (cs *commonService) Start() {
	Test1()
}

func (cs *commonService) RegisterHandlers(typeInt int32, f func(msg *pb.NetMessage)) {
	cs.Lock()
	defer cs.Unlock()

	if _, ok := cs.HandlerMap[typeInt]; !ok {
		cs.HandlerMap[typeInt] = f
	}
}
