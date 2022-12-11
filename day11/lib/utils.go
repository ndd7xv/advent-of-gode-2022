package day11

type Monkey struct {
	items     []int
	operation func(int) int
	test      func(int) int
	inspected int
}

func New() Monkey {
	return Monkey{
		[]int{},
		nil,
		nil,
		0,
	}
}
