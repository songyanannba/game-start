package service

import (
	"fmt"
	"game-player/protoc/pb"
)

func (ps *playerService) PlayerTest1() {
	PlayerService.RegisterHandler(int32(pb.DoType_COMMON), func(ps *playerSpace, msg *pb.NetMessage) {
		fmt.Println("test1 ...")

		m := &pb.GameMessage{
			Name: "看绝代风华开始绝代风华空间",
		}

		NatsManager.SendGateway("syn", m)
	})
}
