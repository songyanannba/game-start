package template

import (
	"elim5/logicPack/base"
	"github.com/samber/lo"
)

func (s *SpinInfo) GetInitWindow() [][]*base.Tag {
	return s.initWindow
}

func (s *SpinInfo) SetInitWindow(list [][]*base.Tag) {
	s.initWindow = list
}

// GetColumnMap 获取列的映射
func (s *SpinInfo) GetColumnMap() map[int][]*base.Tag {
	return lo.GroupBy(s.FindAllTagsQuote(), func(item *base.Tag) int {
		return item.Y
	})
}
