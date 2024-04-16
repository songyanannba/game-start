package src

import (
	"context"
	"fmt"
	"slot6/protoc/pb"
	"slot6/src/component"
	"slot6/src/unit"
)

type SlotService struct {
}

func (s *SlotService) SlotSpin(ctx context.Context, req *pb.SpinReq) (*pb.SpinRes, error) {
	var opts []component.Option
	Opts := append(opts,
		component.SetFreeNum(int(req.FreeNum)),
		component.SetResNum(int(req.ResNum)),
	)

	m, err := unit.Play(uint(req.GameId), int(req.Bet), Opts...)
	if err != nil {
		fmt.Println("Play err", err)
	}

	fmt.Println(m)

	//var spin  *component.Spin
	//var spins []*component.Spin
	//spin = m.GetSpin()
	//spins = m.GetSpins()

	//todo
	spinRes := &pb.SpinRes{
		Msg:  "synsss",
		Code: 0,
		Data: "{name:123 ,str:34}",
	}
	fmt.Println("SlotTest succ and return ...")
	return spinRes, nil
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
