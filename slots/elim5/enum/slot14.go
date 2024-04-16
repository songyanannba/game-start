package enum

const (
	NormalTransferOutWildWeight = 0
	FreeTransferOutWildWeight   = 1
)

var WildMulWeight = map[[2]int][2]int{
	[2]int{SpinAckType1NormalSpin, 1}:   {0, 2},
	[2]int{SpinAckType1NormalSpin, 2}:   {1, 4},
	[2]int{SpinAckType1NormalSpin, 4}:   {2, 8},
	[2]int{SpinAckType1NormalSpin, 8}:   {3, 16},
	[2]int{SpinAckType1NormalSpin, 16}:  {4, 32},
	[2]int{SpinAckType1NormalSpin, 32}:  {5, 64},
	[2]int{SpinAckType1NormalSpin, 64}:  {6, 128},
	[2]int{SpinAckType1NormalSpin, 128}: {7, 256},
	[2]int{SpinAckType1NormalSpin, 256}: {7, 256},
	[2]int{SpinAckType2FreeSpin, 1}:     {8, 3},
	[2]int{SpinAckType2FreeSpin, 3}:     {9, 9},
	[2]int{SpinAckType2FreeSpin, 9}:     {10, 27},
	[2]int{SpinAckType2FreeSpin, 27}:    {11, 81},
	[2]int{SpinAckType2FreeSpin, 81}:    {12, 243},
	[2]int{SpinAckType2FreeSpin, 243}:   {13, 729},
	[2]int{SpinAckType2FreeSpin, 729}:   {13, 729},
}
