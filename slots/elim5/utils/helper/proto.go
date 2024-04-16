package helper

import (
	"elim5/pbs/common"
	"google.golang.org/protobuf/proto"
)

func IdsToProto(ids []uint64) []byte {
	buf, _ := proto.Marshal(&common.Ids{Ids: ids})
	return buf
}

func ProtoToIds(ids []byte) []uint64 {
	if len(ids) == 0 {
		return []uint64{}
	}
	var idArr common.Ids
	_ = proto.Unmarshal(ids, &idArr)
	return idArr.Ids
}

//func ProtoToSurvival(survival []byte) map[int32]*common.SurvivalDayInfo {
//	var survivalp common.Survival
//	_ = proto.Unmarshal(survival, &survivalp)
//	return survivalp.SurvivalDay
//}

func BkPlayerToProto(bkPlayer map[uint64]int64) []byte {
	buf, _ := proto.Marshal(&common.BkPlayer{PlayerBkMap: bkPlayer})
	return buf
}

func ProtoToBkPlayer(bkPlayer []byte) map[uint64]int64 {
	if len(bkPlayer) == 0 {
		return map[uint64]int64{}
	}
	var bkPlayerP common.BkPlayer
	_ = proto.Unmarshal(bkPlayer, &bkPlayerP)
	return bkPlayerP.PlayerBkMap
}

func SurvivalToProto(survival *common.SurvivalMap) []byte {
	buf, _ := proto.Marshal(survival)
	return buf
}

func ProtoToSurvivalMap(survival []byte) *common.SurvivalMap {
	if len(survival) == 0 {
		return &common.SurvivalMap{
			DaySurvival: make(map[uint32]*common.SurvivalData),
		}
	}
	var survivalP common.SurvivalMap
	_ = proto.Unmarshal(survival, &survivalP)
	return &survivalP
}
