package cache

import (
	"elim5/model/business"
	"elim5/pbs/common"
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/samber/lo"
	"sort"
	"sync/atomic"
)

var (
	SlotCache   = cmap.NewWithCustomShardingFunction[uint, *business.Slot](sharding)
	SlotCacheOk = atomic.Bool{}
)

func checkSlotCache() error {
	if SlotCacheOk.Load() {
		return nil
	}

	slots, err := business.GetList[*business.Slot]()
	if err != nil {
		return err
	}

	slotMap := lo.SliceToMap(slots, func(item *business.Slot) (k uint, v *business.Slot) {
		return item.ID, item
	})

	var fileMap map[uint][]*common.File
	fileMap, err = GetFileMapByCache()
	if err != nil {
		return err
	}

	for _, slot := range slotMap {
		slot.FileList = fileMap[slot.ID]
	}

	SlotCache.MSet(slotMap)
	SlotCacheOk.Store(true)
	return nil
}

func GetSlot(id uint) (res *business.Slot, err error) {
	if err = checkSlotCache(); err != nil {
		return nil, err
	}
	var ok bool
	res, ok = SlotCache.Get(id)
	if ok {
		return res, nil
	}

	return res, fmt.Errorf("id: %d not found", id)
}

func GetSlotList() (res []*business.Slot, err error) {
	if err = checkSlotCache(); err != nil {
		return nil, err
	}
	for _, v := range SlotCache.Items() {
		res = append(res, v)
	}
	return res, nil
}

func GetSlotMap() (m map[uint]*business.Slot, err error) {
	if err = checkSlotCache(); err != nil {
		return nil, err
	}
	return SlotCache.Items(), nil
}

func ClearSlotCache() {
	SlotCacheOk.Store(false)
	SlotCache.Clear()
}

func GetGameList(needConf bool) []*common.GameInfo {
	var (
		slots, _ = GetSlotList()
		gameList []*common.GameInfo
	)
	for _, slot := range slots {
		game := &common.GameInfo{
			Id:     int32(slot.ID),
			Name:   slot.NamePkg,
			Icon:   slot.Icon,
			Url:    slot.Url,
			Status: int32(slot.Status),
		}
		if needConf {
			game.Config = slot.ClientConf
			game.FileList = slot.FileList
		}
		gameList = append(gameList, game)
	}
	sort.Slice(gameList, func(i, j int) bool {
		return gameList[i].Id < gameList[j].Id
	})
	return gameList
}
