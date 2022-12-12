package day11

// Monkey represents the item flinging devils as described in the problem.
//
// items represents an array of items, identified by the worry value they incur
// operation is the closure that updates the worry value when inspected by a monkey
// test is the closure that determines what monkey to throw it to next
// inspected is an incrementor that is increased every time the monkey inspects an item
type Monkey struct {
	items     []int
	operation func(int) int
	test      func(int) int
	inspected int
}

// New creates a new Monkey.
//
// the operation and test closures are determined when the input is parsed.
func New() Monkey {
	return Monkey{
		[]int{},
		nil,
		nil,
		0,
	}
}
