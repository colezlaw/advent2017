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

var bits = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'a': "1010",
	'b': "1011",
	'c': "1100",
	'd': "1101",
	'e': "1110",
	'f': "1111",
}

func makeRow(s string) []int {
	res := make([]int, 128)
	for b, r := range s {
		for c, n := range bits[r] {
			if n == '1' {
				res[b*4+c] = 1
			}
		}
	}
	return res
}

func regionize(grid, target [][]int, r, c, maxregion int) {
	if r > 0 && grid[r-1][c] == 1 && target[r-1][c] == 0 {
		target[r-1][c] = maxregion
		regionize(grid, target, r-1, c, maxregion)
	}
	if c > 0 && grid[r][c-1] == 1 && target[r][c-1] == 0 {
		target[r][c-1] = maxregion
		regionize(grid, target, r, c-1, maxregion)
	}
	if r < len(grid)-1 && grid[r+1][c] == 1 && target[r+1][c] == 0 {
		target[r+1][c] = maxregion
		regionize(grid, target, r+1, c, maxregion)
	}
	if c < len(grid[r])-1 && grid[r][c+1] == 1 && target[r][c+1] == 0 {
		target[r][c+1] = maxregion
		regionize(grid, target, r, c+1, maxregion)
	}
}

func main() {
	key := "amgozmfv"
	grid := make([][]int, 128)
	for i := 0; i < 128; i++ {
		c := newChain(256, []byte(fmt.Sprintf("%s-%d", key, i)))
		str := c.hash()
		grid[i] = makeRow(str)
	}

	target := make([][]int, 128)
	for r := 0; r < 128; r++ {
		target[r] = make([]int, 128)
	}

	// Find all the 1's in the hash grid and calculate regions in the target grid
	maxregion := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 1 && target[r][c] == 0 {
				maxregion++
				regionize(grid, target, r, c, maxregion)
			}
		}
	}

	for _, r := range target {
		for _, c := range r {
			fmt.Printf("%5d", c)
		}
		fmt.Println()
	}

	fmt.Println(maxregion)
}
