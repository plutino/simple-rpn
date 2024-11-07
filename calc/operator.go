package calc

type operator interface {
	Do(int, int) int
}

type add struct{}
type sub struct{}
type mul struct{}
type div struct{}
type pow struct{}

func (add) Do(a, b int) int {
	return a + b
}

func (sub) Do(a, b int) int {
	return a - b
}

func (mul) Do(a, b int) int {
	return a * b
}

func (div) Do(a, b int) int {
	return a / b
}

func (pow) Do(a, b int) int {
	res := 1
	for i := 0; i < b; i++ {
		res *= a
	}
	return res
}
