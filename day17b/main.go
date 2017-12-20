package main

import "fmt"

type spinlock struct {
	cbuf      int
	index     int
	count     int
	steps     int
	slot1     int
	zeroindex int
}

func newSpinlock(steps int) *spinlock {
	return &spinlock{cbuf: 1, steps: steps}
}

func (s *spinlock) spin() {
	s.count++
	s.index = (s.index+s.steps)%s.cbuf + 1
	if s.index == 0 {
		// We're inserting before element 0, so we shift the 0 index
		s.zeroindex++
	}
	if s.index == s.zeroindex+1 {
		fmt.Println(s.count)
		s.slot1 = s.count
	}
	if s.index == 0 {
		fmt.Println("yikes!")
	}
	s.cbuf++
}

func main() {

}
