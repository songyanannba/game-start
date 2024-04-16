package unit

import (
	"elim5/logicPack/component"
)

type SlotFace interface {
	RunTem() ([]*component.Spin, error)
	Calculate(spins []*component.Spin)
	GetStatus() (float64, string, bool)
}

type SlotFaceImp struct {
}

func (si *SlotFaceImp) RunTem() ([]*component.Spin, error) {

	return nil, nil
}

func (si *SlotFaceImp) Calculate(spins []*component.Spin) {

}

func (si *SlotFaceImp) GetStatus() (float64, string, bool) {
	return 0, "", false
}
