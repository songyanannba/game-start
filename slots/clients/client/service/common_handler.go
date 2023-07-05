package service

import "client/protoc/pb"

func Test1() {
	CommonService.RegisterHandlers(int32(pb.DoType_DI_YI), func(msg *pb.NetMessage) {
		request := &pb.GameMessage{}
		request.Do = "sd"
		request.To = "hhahah"
		request.Todo = "会面吧"

		//发送

	})
}
