package src

import (
	"context"
	"elim5/logicPack"
	"elim5/logicPack/component"
	"elim5/pbs/common"
	"fmt"
	"google.golang.org/protobuf/proto"
)

// ElimService
type ElimService struct {
}

func (s *ElimService) SlotTest(ctx context.Context, req *common.SpinReq) (*common.SpinRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ElimService) SlotSpin(ctx context.Context, req *common.SpinReq) (*common.SpinRes, error) {
	var opts []component.Option
	Opts := append(opts,
		component.SetFreeNum(int(req.FreeNum)),
		component.SetResNum(int(req.ResNum)),
	)

	m, err := logicPack.Play(uint(req.GameId), int(req.Bet), Opts...)
	if err != nil {
		fmt.Println("Play err", err)
	}

	fmt.Println(m)

	//var spin *component.Spin
	//var spins []*component.Spin
	//spin = m.GetSpin()
	//spins = m.GetSpins()
	//fmt.Println(spin, spins)

	sps := append([]*component.Spin{m.GetSpin()}, m.GetSpins()...)
	ack := SumAck(sps)

	ackMarshal, _ := proto.Marshal(ack)

	//todo
	spinRes := &common.SpinRes{
		Msg:  "成功",
		Code: 0,
		MsgData: &common.NetMessage{
			ServiceId: "",
			UId:       "",
			Content:   ackMarshal,
			Type:      0,
			SlotID:    5,
		},
	}
	fmt.Println("SlotTest succ and return ...")
	return spinRes, nil
}

func SumAck(spins []*component.Spin) *common.MatchSpinAck {
	ack := &common.MatchSpinAck{
		Steps: make([]*common.MatchSpinStep, 0),
	}

	for _, spin := range spins {
		step := &common.MatchSpinStep{
			Id:         int32(spin.Id),
			Pid:        int32(spin.ParentId),
			Type:       int32(spin.Type()),
			SumGain:    int64(spin.Gain),
			InitList:   spin.SpinInfo.GetInitAck(),
			Flows:      make([]*common.StepFlow, 0),
			SingleTags: make([]*common.SingleTag, 0),
			TemInit:    spin.SpinInfo.GetTemInitRows(),
			OtherLine:  []*common.Tags{spin.SpinInfo.Scatter.ToAck(1)},
		}
		for _, flow := range spin.SpinInfo.SpinFlow {
			step.Flows = append(step.Flows, flow.ToAck(1))
		}
		step.SingleTags = spin.SpinInfo.GetSingleTags() //wild轨迹
		ack.Steps = append(ack.Steps, step)
	}

	return ack
}
