package main

import "fmt"

type chain struct {
	list    []int
	pos     int
	skip    int
	lengths []int
}

func newChain(size int, lengths []int) *chain {
	c := &chain{list: make([]int, size), lengths: lengths}

	for i := 0; i < size; i++ {
		c.list[i] = i
	}

	return c
}

func (c *chain) round() {
	reverse(c.list, c.pos, c.lengths[c.skip])
	c.pos = (c.pos + c.lengths[c.skip] + c.skip) % len(c.list)
	c.skip++
}

func reverse(input []int, pos, l int) {
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		input[(pos+i)%len(input)], input[(pos+j)%len(input)] = input[(pos+j)%len(input)], input[(pos+i)%len(input)]
	}
}

func (c *chain) hash() {
	for c.skip < len(c.lengths) {
		c.round()
	}
}

func main() {
	c := newChain(256, []int{120, 93, 0, 90, 5, 80, 129, 74, 1, 165, 204, 255, 254, 2, 50, 113})
	c.hash()
	fmt.Printf("%d\n", c.list[0]*c.list[1])
}
