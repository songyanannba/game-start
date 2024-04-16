package logic

import (
	"elim5/logicPack/base"
	"fmt"
)

func (d *Dismantling) GetGameType() int {
	return int(d.TemGens[0].Type)
}

func (d *Dismantling) GetRtp() int {
	return d.TemGens[0].Rtp
}

func (d *Dismantling) GetTemTag(x, y int) *base.Tag {
	nowX := x
	for {
		if nowX >= len(d.Template[y]) {
			nowX = nowX - len(d.Template[y])
		} else {
			break
		}
	}

	fillTag := d.Template[y][nowX]

	return fillTag.Copy()
}

func (d *Dismantling) GetTemLayout(y int) string {
	temStr := ""
	for _, tag := range d.Template[y] {
		temStr += fmt.Sprintf("%d,", tag.Id)
	}
	return temStr
}

func (d *Dismantling) GetTemResLayout(y int) string {
	temStr := ""
	for _, tag := range d.TemplateRes[y] {
		temStr += fmt.Sprintf("%d,", tag)
	}
	return temStr
}
