package unit

import "slot6/src/component"

type MachineItf interface {
	GetSpin() *component.Spin
	Exec()
	GetInitData()
	GetResData()
	SumGain()
	GetSpins() []*component.Spin
}

func RunSpin(s *component.Spin) (m MachineItf, err error) {
	m = NewMachine(s)
	m.Exec()
	return m, nil
}

func Play(slotId uint, amount int, options ...component.Option) (m MachineItf, err error) {
	var s *component.Spin
	s, err = component.NewSpin(slotId, amount, options...)
	if err != nil {
		return nil, err
	}
	return RunSpin(s)
}
