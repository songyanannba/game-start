package merger

import (
	businessReq "elim5/model/business/request"
	"elim5/template_gen/merger/logic"
)

var isMerger = false

// GetIsMerger 获取是否正在合并
func GetIsMerger() bool {
	return isMerger
}

// TemMerger 模板合并
func TemMerger(req *businessReq.MergerIncrease) (err error) {
	isMerger = true
	defer func() {
		isMerger = false
		errStr := "合并成功"
		if err != nil {
			errStr = err.Error()
		}
		logic.WritingProgress(0, 0, errStr)
	}()

	var dismantling *logic.Dismantling
	dismantling, err = logic.NewDismantling(req)
	if err != nil {
		return err
	}
	err = dismantling.Merger()
	if err != nil {
		return err
	}
	err = dismantling.Save()
	return err
}
