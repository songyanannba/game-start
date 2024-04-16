package eliminate

import (
	"elim5/logicPack/base"
)

func (t *Table) EliminatedTagNameSetEmpty(tagList [][]*base.Tag) [][]*base.Tag {
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
