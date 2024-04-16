package cache

//func GetUserSlot21AckKey(userId int64, slotId int) string {
//	return fmt.Sprintf("{user_slot_ack}:%s:%s", userId, slotId)
//}
//
//func SetUserSlot21Ack(userId int64, slotId int, ack *game21.SpinAck) (err error) {
//	key := GetUserSlot21AckKey(userId, slotId)
//	var data []byte
//	data, err = proto.Marshal(ack)
//	if err != nil {
//		return
//	}
//	err = global.GVA_REDIS.Set(context.Background(), key, data, time.Hour*24*15).Err()
//	return
//}
//
//func GetUserSlot21Ack(userId int64, slotId int) (ack *game21.SpinAck, err error) {
//	ack = &game21.SpinAck{}
//	key := GetUserSlot21AckKey(userId, slotId)
//	var result []byte
//	result, err = global.GVA_REDIS.Get(context.Background(), key).Bytes()
//	if err != nil {
//		return
//	}
//	err = proto.Unmarshal(result, ack)
//	return
//}
//
//func DelUserSlot21Ack(userId int64, slotId int) (err error) {
//	return global.GVA_REDIS.Del(context.Background(), GetUserSlot21AckKey(userId, slotId)).Err()
//}

//
//func GetUserRecordIdKey(userId int64) string {
//	return "{user_rcd_id}:" + strconv.Itoa(int(userId))
//}
//
//func GetUserRecordId(userId int64) (sRcd *business.SlotRecord, err error) {
//	sRcd = &business.SlotRecord{}
//	key := GetUserRecordIdKey(userId)
//	var result []byte
//	result, err = global.GVA_REDIS.Get(context.Background(), key).Bytes()
//	if err != nil {
//		return
//	}
//	err = global.Json.Unmarshal(result, sRcd)
//	return
//}
//
//func SetUserRecordId(userId int64, sRcd business.SlotRecord) (err error) {
//	key := GetUserRecordIdKey(userId)
//	var data []byte
//	data, err = global.Json.Marshal(sRcd)
//	if err != nil {
//		return
//	}
//	err = global.GVA_REDIS.Set(context.Background(), key, data, time.Hour*24*15).Err()
//	return
//}
//
//func DelUserRecordId(userId int64) (err error) {
//	return global.GVA_REDIS.Del(context.Background(), GetUserRecordIdKey(userId)).Err()
//}
