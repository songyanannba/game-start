package logic

type Eliminate interface {
	Merger(index int) error
}
