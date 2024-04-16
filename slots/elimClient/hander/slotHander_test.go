package hander

import (
	"testing"
)

func TestSlot6SpinTest(t *testing.T) {
	Slot6SpinTest()
}

func BenchmarkSlot6SpinTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slot6SpinTest()
	}
}
