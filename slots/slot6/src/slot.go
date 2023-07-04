package src

import (
	"context"
	"fmt"
	"slot6/protoc/pb"
)

type SlotService struct {
}

func (s *SlotService) SlotTest(ctx context.Context, req *pb.SlotTestReq) (*pb.SlotTestRes, error) {

	fmt.Println("SlotTestReq === ", req)

	slotTestRes := &pb.SlotTestRes{
		Msg:  "synsss",
		Code: 0,
		Data: "{name:123 ,str:34}",
	}
	fmt.Println("SlotTest succ and return ...")
	return slotTestRes, nil
}

func (s *SlotService) Start() {

	//afterTime := time.After(time.Second * 3)
	//
	//for {
	//	select {
	//	case <-afterTime:
	//		fmt.Println("afterTime 3..")
	//	}
	//}
}
