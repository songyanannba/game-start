package template

import (
	"elim5/logicPack/template/flow"
)

// AlterDisplay 修改指定位置标签
func (s *SpinInfo) AlterDisplay(list []*flow.WinLine) {
	for _, tags := range list {
		for _, tag := range tags.Tags {
			s.Display[tag.X][tag.Y] = tag.Copy()
		}
	}
}
