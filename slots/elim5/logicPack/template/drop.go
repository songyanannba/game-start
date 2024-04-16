package template

import (
	"elim5/logicPack/base"
)

//第八台游戏逻辑

// DropDisplay 当前窗口空白位置上浮
func (s *SpinInfo) DropDisplay() {
	tags := s.GetEmptyTags() //Display
	for _, tag := range tags {
		for x := tag.X; x > 0; x-- {
			if len(s.Display[x-1]) <= tag.Y {
				continue
			}
			if s.Display[x-1][tag.Y].IsFixed() {
				continue
			}
			//和上面的交换位置
			s.Display[x][tag.Y], s.Display[x-1][tag.Y] = s.Display[x-1][tag.Y], s.Display[x][tag.Y]
		}
	}
	s.SetAllLocation()
}

func (s *SpinInfo) Drop() []*base.Tag {
	s.DropDisplay()
	tags := make([]*base.Tag, 0)
	for x := len(s.Display) - 1; x >= 0; x-- {
		row := s.Display[x]
		for y, tag := range row {
			if tag.IsEmpty() {
				fillTag := s.GetTemplateTag(x, y)
				s.Display[x][y] = fillTag //列向下转
				tags = append(tags, fillTag.Copy())
			}
		}
	}
	return tags
}

func (s *SpinInfo) DropAndFill(tags []*base.Tag) (string, []*base.Tag) {
	for _, tag := range tags {
		for i := 0; i < tag.X; i++ {
			if s.Display[i][tag.Y].IsWild || s.Display[i][tag.Y].Name == "" {
				continue
			}
			s.Display[tag.X][tag.Y], s.Display[i][tag.Y] = s.Display[i][tag.Y].Copy(), s.Display[tag.X][tag.Y].Copy()
		}
	}
	s.SetAllLocation()
	//
	str := s.PrintTable("掉落")

	addTags := make([]*base.Tag, 0)
	for x := len(s.Display) - 1; x >= 0; x-- {
		row := s.Display[x]
		for y, tag := range row {
			if tag.Name == "" {
				fillTag := s.GetTemplateTag(x, y)
				s.Display[x][y] = fillTag //列向下转
				addTags = append(addTags, fillTag.Copy())
			}
		}
	}

	return str, addTags
}

//func (s *AlienInfo) DropFor9() (string, []*base.Tag) {
//	tags := s.GetEmptyTags()
//	//sort.Slice(tags, func(i, j int) bool {
//	//	return tags[i].X < tags[j].X
//	//})
//	for _, tag := range tags {
//		for i := 0; i < tag.X; i++ {
//			if s.Display[i][tag.Y].IsWild || s.Display[i][tag.Y].Name == "" {
//				continue
//			}
//			s.Display[tag.X][tag.Y], s.Display[i][tag.Y] = s.Display[i][tag.Y].Copy(), s.Display[tag.X][tag.Y].Copy()
//		}
//	}
//	s.SetAllLocation()
//	//
//	str := s.PrintTable(enum.Drop)
//
//	addTags := make([]*base.Tag, 0)
//	for x := len(s.Display) - 1; x >= 0; x-- {
//		row := s.Display[x]
//		for y, tag := range row {
//			if tag.Name == "" {
//				fillTag := s.GetTemplateTag(x, y)
//				s.Display[x][y] = fillTag //列向下转
//				addTags = append(addTags, fillTag.Copy())
//			}
//		}
//	}
//
//	return str, addTags
//}
