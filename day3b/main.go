package main

import (
	"fmt"
	"math"
)

// Get the next direction we should try to go
func nextDir(dir int) int {
	return (dir + 1) % 4
}

// Get the cell coordinates to the specified direction
func coord(y, x, dir int) (int, int) {
	fmt.Printf("dir is %d\n", dir)
	switch dir {
	case 0:
		return y + 1, x
	case 1:
		return y, x + 1
	case 2:
		return y - 1, x
	}
	return y, x - 1
}

func printGrid(foo [][]int) {
	for y := 0; y < len(foo); y++ {
		for x := 0; x < len(foo[y]); x++ {
			fmt.Printf("%4d", foo[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
}

// Figure out the lowest neighbor (non-diagonal)
func nextStep(grid [][]int, y, x int) (minY, minX int) {
	min := grid[y][x]
	if min == 1 {
		return y, x
	}
	// Look to the right
	if x < len(grid[y])-1 {
		if grid[y][x+1] > 0 && grid[y][x+1] < min {
			min, minY, minX = grid[y][x+1], y, x+1
		}
	}
	// Look below
	if y < len(grid)-1 {
		if grid[y+1][x] > 0 && grid[y+1][x] < min {
			min, minY, minX = grid[y+1][x], y+1, x
		}
	}
	// Look to the left
	if x > 0 {
		if grid[y][x-1] > 0 && grid[y][x-1] < min {
			min, minY, minX = grid[y][x-1], y, x-1
		}
	}
	// Look up
	if y > 0 {
		if grid[y-1][x] != 0 && grid[y-1][x] < min {
			min, minY, minX = grid[y-1][x], y-1, x
		}
	}
	return minY, minX
}

func sum(grid [][]int, x, y int) int {
	sum := 0
	// Get the cells above
	if y > 0 {
		if x > 0 {
			sum += grid[y-1][x-1]
		}
		sum += grid[y-1][x]
		if x < len(grid[y-1])-1 {
			sum += grid[y-1][x+1]
		}
	}
	// Get the cells on the same row
	if x > 0 {
		sum += grid[y][x-1]
	}
	if x < len(grid[y])-1 {
		sum += grid[y][x+1]
	}
	// Get the cells below
	if y < len(grid)-1 {
		if x > 0 {
			sum += grid[y+1][x-1]
		}
		sum += grid[y+1][x]
		if x < len(grid[y+1])-1 {
			sum += grid[y+1][x+1]
		}
	}

	return sum
}

func main() {
	target := 325489
	size := int(math.Sqrt(float64(target)) + 2.0)
	fmt.Println(size)
	foo := make([][]int, size)
	for i := 0; i < size; i++ {
		foo[i] = make([]int, size)
	}
	x := size/2 - 1
	y := size/2 - 1
	dir := 3
	foo[x][y] = 1
	for n := 1; n <= target; {
		// Get the next direction we should go, and peek at that cell
		newDir := nextDir(dir)
		newY, newX := coord(y, x, newDir)
		if foo[newY][newX] == 0 {
			// We can turn that direction
			dir = newDir
		}
		y, x = coord(y, x, dir)
		n = sum(foo, x, y)
		fmt.Printf("Setting {%d,%d} to %d\n", x, y, n)
		foo[y][x] = n
	}

	count := 0
	for test := target; test != 1; y, x = nextStep(foo, y, x) {
		test = foo[y][x]
		if test == 1 {
			break
		}
		count++
	}
	fmt.Println(count)
}
