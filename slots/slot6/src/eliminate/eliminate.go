package eliminate

import (
	"slot6/src/base"
)

func (t *Table) EliminateV2(tagList [][]*base.Tag) [][]*base.Tag {
	for _, tags := range tagList {
		for _, tag := range tags {
			if !t.TagList[tag.X][tag.Y].IsWild {
				t.TagList[tag.X][tag.Y] = &base.Tag{
					Name: "",
					Y:    tag.Y,
					X:    tag.X,
				}
			}
		}
	}
	return tagList
}
