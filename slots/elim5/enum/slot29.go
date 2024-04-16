package enum

const (
	Slot29TagCountWeight = iota
	Slot29TagWeight
	Slot29TagMulCount
)

const TakeOffer = "takeOffer"
const TryAgain = "tryAgain"

// multiple_5_1、multiple_5_2
// multiple_10_1、multiple_10_2
// multiple_20_1、multiple_20_2
// multiple_50_1、multiple_50_2
// multiple_100
// multiple_1000
// multiple_x2_1、multiple_x2_2
const (
	Multiple51   = "5_1"
	Multiple52   = "5_2"
	Multiple101  = "10_1"
	Multiple102  = "10_2"
	Multiple201  = "20_1"
	Multiple202  = "20_2"
	Multiple501  = "50_1"
	Multiple502  = "50_2"
	Multiple100  = "100"
	Multiple1000 = "1000"
	Multiplex21  = "x2_1"
	Multiplex22  = "x2_2"
)

var MultipleMap = map[string]int{
	Multiple51:   5,
	Multiple52:   5,
	Multiple101:  10,
	Multiple102:  10,
	Multiple201:  20,
	Multiple202:  20,
	Multiple501:  50,
	Multiple502:  50,
	Multiple100:  100,
	Multiple1000: 1000,
	Multiplex21:  2,
	Multiplex22:  2,
}
