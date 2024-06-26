package template

import (
	"elim5/logicPack/base"
)

// DeleteTag 删除指定位置的tag
func (s *SpinInfo) DeleteTag(tag *base.Tag) {
	s.Display[tag.X][tag.Y] = &base.Tag{
		Name:     "",
		X:        tag.X,
		Y:        tag.Y,
		Multiple: 1,
	}
}

// DeleteTags 删除指定位置的一组tag
func (s *SpinInfo) DeleteTags(tags []*base.Tag) {
	for _, tag := range tags {
		s.DeleteTag(tag)
	}
}

// DeleteTagList 删除指定位置的多组tag
func (s *SpinInfo) DeleteTagList(list [][]*base.Tag) {
	for _, tags := range list {
		s.DeleteTags(tags)
	}
}
