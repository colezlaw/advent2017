package main

import "fmt"

func maxindex(banks []int) int {
	if len(banks) == 0 {
		return 0
	}
	max := 0
	// It must return the first highest index when multiple
	// share the same highest value
	for k, v := range banks {
		if v > banks[max] {
			max = k
		}
	}

	return max
}

func distribute(banks []int) {
	if len(banks) == 0 {
		return
	}

	// Get the bank with the most work
	start := maxindex(banks)
	index := start

	// Set that bank to 0
	count := banks[start]
	banks[start] = 0

	for count > 0 {
		index = (index + 1) % len(banks)
		banks[index]++
		count--
	}
}

func optimize(banks []int) int {
	seen := make(map[string]bool)
	count := 0
	seen[fmt.Sprintf("%v", banks)] = true

	for {
		count++
		distribute(banks)
		if _, ok := seen[fmt.Sprintf("%v", banks)]; ok {
			break
		}
		seen[fmt.Sprintf("%v", banks)] = true
	}

	return count
}

func main() {
	banks := []int{0, 5, 10, 0, 11, 14, 13, 4, 11, 8, 8, 7, 1, 4, 12, 11}
	count := optimize(banks)
	fmt.Printf("%d %v\n", count, banks)
}
