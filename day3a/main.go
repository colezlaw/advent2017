package main

func nextDir(dir int) int {
	return (dir + 1) % 4
}

func nextCell(grid *[][]int, dir, row, col int) (newDir, newR, newC int) {
	newDir := nextDir(dir)
	switch newDir {
	case 0:
		newR := r - 1
	case 1:
		newC := c + 1
	case 2:
		newR := r + 1
	case 3:
		newC := c + 1
	}
	// See if we can change direction first
	if newR < 0 {
		grid = append([][]int, grid...)
	}

	// If not, move one more cell the direction we're going
}

func main() {
	max := 23
	dir := 0
	grid := make([][]int, 1)
	grid[0] = make([]int, 1)
	grid[0][0] = 1
}
