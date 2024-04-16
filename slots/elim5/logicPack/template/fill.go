package template

import (
	"elim5/enum"
	"elim5/global"
	"elim5/logicPack/base"
	"elim5/utils/helper"
)

func (s *SpinInfo) FillInit() error {
	s.SetAllLocation()

	isBuy := s.Config.GetIsBuy()
	if s.infoType == enum.SpinAckType1NormalSpin && isBuy != enum.NoBuy {
		err := s.FillCustomDisplay(isBuy)
		if err != nil {
			return err
		}
	} else {
		s.FillInitDisplay()
	}
	s.initWindow = s.InitConvert(s)
	return nil
}

// FillInitDisplay 填充初始展示
func (s *SpinInfo) FillInitDisplay() {
	for i := len(s.Display) - 1; i >= 0; i-- {
		for i2, _ := range s.Display[i] {
			if s.Display[i][i2].IsEmpty() {
				fillTag := s.GetTemplateTag(i, i2)
				if fillTag == nil {
					global.GVA_LOG.Error("fillTag==nil")
				}
				s.Display[i][i2] = fillTag
			}
		}
	}
}

// FillCustomDisplay 自定义填充函数
func (s *SpinInfo) FillCustomDisplay(isBuy int) error {
	if s.CustomFill != nil {
		err := s.CustomFill[isBuy](s)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *SpinInfo) FillAssignScat() {
	s.FillCustomDisplay(enum.SpinAckType2FreeSpin)
	s.initWindow = helper.CopyListArr(s.Display)
}

func (s *SpinInfo) FillNoLine(number int) {
	fillTags := s.GetNormalTags()
	emptyTags := s.GetEmptyTags()
	for _, tag := range emptyTags {
		for {
			fillTag := fillTags[helper.RandInt(len(fillTags))].Copy()
			fillTag.X = tag.X
			fillTag.Y = tag.Y
			s.Display[tag.X][tag.Y] = fillTag
			if len(s.FindAdjacentLine(number)) > 0 {
				continue
			} else {
				break
			}
		}
	}
}

func (s *SpinInfo) FillCountNoLine() {
	fillTags := s.GetNormalTags()
	emptyTags := s.GetEmptyTags()
	for _, tag := range emptyTags {
		for {
			fillTag := fillTags[helper.RandInt(len(fillTags))].Copy()
			fillTag.X = tag.X
			fillTag.Y = tag.Y
			s.Display[tag.X][tag.Y] = fillTag
			if len(s.FindSpecifyCountLine(fillTag, 8)) > 0 {
				continue
			} else {
				break
			}
		}
	}
}

func (s *SpinInfo) FillTags(tags []*base.Tag) {
	for _, tag := range tags {
		s.Display[tag.X][tag.Y] = tag.Copy()
	}
}
