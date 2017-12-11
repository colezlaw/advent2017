package main

import "fmt"

type chain struct {
	list    []byte
	pos     int
	skip    int
	lengths []byte
}

func newChain(size int, lengths []byte) *chain {
	c := &chain{list: make([]byte, size), lengths: lengths}

	for i := 0; i < size; i++ {
		c.list[i] = byte(i)
	}

	c.lengths = append(c.lengths, 17, 31, 73, 47, 23)

	return c
}

func (c *chain) round() {
	for i := 0; i < len(c.lengths); i++ {
		reverse(c.list, c.pos, c.lengths[i])
		c.pos = (c.pos + int(c.lengths[i]) + c.skip) % len(c.list)
		c.skip++
	}
}

func reverse(input []byte, pos int, l byte) {
	for i, j := byte(0), l-1; i < j; i, j = i+1, j-1 {
		input[(pos+int(i))%len(input)], input[(pos+int(j))%len(input)] = input[(pos+int(j))%len(input)], input[(pos+int(i))%len(input)]
	}
}

func (c *chain) hash() string {
	for i := 0; i < 64; i++ {
		c.round()
	}

	dense := make([]byte, 16)
	for block := 0; block < len(dense); block++ {
		acc := c.list[block*16]
		for i := 1; i < 16; i++ {
			acc = acc ^ c.list[block*16+i]
		}
		dense[block] = acc
	}

	return fmt.Sprintf("%x", dense)
}

func main() {
	c := newChain(256, []byte("120,93,0,90,5,80,129,74,1,165,204,255,254,2,50,113"))
	fmt.Println(c.hash())
}
