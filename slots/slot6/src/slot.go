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

type Machine interface {
	GetSpin() *component.Spin
	Exec()
	GetInitData()
	GetResData()
	SumGain()
	GetSpins() []*component.Spin
}

func RunSpin(s *component.Spin) (m Machine, err error) {
	m = unit.NewMachine(s)
	m.Exec()
	return m, nil
}

func Play(slotId uint, amount int, options ...component.Option) (m Machine, err error) {
	var s *component.Spin
	s, err = component.NewSpin(slotId, amount, options...)
	if err != nil {
		return nil, err
	}
	return RunSpin(s)
}

func (s *SlotService) SlotSpin(ctx context.Context, req *pb.SpinReq) (*pb.SpinRes, error) {
	var opts []component.Option
	Opts := append(opts,
		component.SetFreeNum(int(req.FreeNum)),
		component.SetResNum(int(req.ResNum)),
	)

	Play(uint(req.GameId), int(req.Bet), Opts...)

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
