package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// For each adjacent pair of digits in the input, if they match, accumulate
// the digit as a sum
func findSum(input []byte) int {
	sum := 0

	// General case: first item compared to second item, up to last-1 compared to last
	for i, j := 0, 1; i < len(input)-1; i, j = i+1, j+1 {
		// Verify it's a digit
		if input[i]-0x30 < 0 || input[i]-0x30 > 9 || input[j]-0x30 < 0 || input[j]-0x30 > 9 {
			continue
		}
		if input[i] == input[j] {
			sum += int(input[i] - 0x30)
		}
	}

	// Special case: compare the last item to the first item
	if input[len(input)-1] == input[0] {
		sum += int(input[0] - 0x30)
	}

	return sum
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("cannot read: %v", err)
	}

	sum := findSum(input)
	fmt.Println(sum)
}
