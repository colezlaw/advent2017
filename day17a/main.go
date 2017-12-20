package main

type spinlock struct {
	cbuf  []int
	index int
	count int
	steps int
}

func newSpinlock(steps int) *spinlock {
	return &spinlock{cbuf: []int{0}, steps: steps}
}

func (s *spinlock) spin() {
	s.count++
	nbuf := make([]int, len(s.cbuf)+1)
	s.index = (s.index + s.steps + 1) % len(s.cbuf)
	if s.index > 0 {
		copy(nbuf, s.cbuf[0:s.index])
	}
	nbuf[s.index] = s.count
	copy(nbuf[s.index+1:], s.cbuf[s.index:])
	s.cbuf = nbuf
}

func main() {

}
