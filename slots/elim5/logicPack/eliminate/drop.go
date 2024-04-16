package eliminate

import (
	"elim5/logicPack/base"
	"sort"
)

func (t *Table) Drop() [][]*base.Tag {
	tags := t.NeedFill()

	sort.Slice(tags, func(i, j int) bool {
		return tags[i].X < tags[j].X
	})

	for _, i2 := range tags {
		for i := 0; i < i2.X; i++ {
			if t.TagList[i][i2.Y].IsWild || t.TagList[i][i2.Y].Name == "" {
				continue
			}
			t.TagList[i2.X][i2.Y], t.TagList[i][i2.Y] = t.TagList[i][i2.Y].Copy(), t.TagList[i2.X][i2.Y].Copy()

		}
	}
	t.SetCoordinates()
	return t.GetGraph()
}
