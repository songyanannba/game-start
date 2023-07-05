package test

import (
	"slot6/core"
	"slot6/src/unit"
	"testing"
)

func TestSlot6(t *testing.T) {

}

func BenchmarkSlot6(b *testing.B) {
	core.BaseInit()
	for i := 0; i < b.N; i++ {
		unit.Play(6, 100)
	}
}
