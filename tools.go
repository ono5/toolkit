// tools.go
package toolkit

type Tools struct {
	sum int
}

func (t *Tools) Increment(n int) int {
	t.sum += n
	return t.sum
}

func (t *Tools) Decrement(n int) int {
	t.sum -= n
	return t.sum
}
