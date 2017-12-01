package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func findSum(input []byte) int {
	sum := 0
	for i, j := 0, len(input)/2; i < len(input); i, j = i+1, (j+1)%len(input) {
		if input[i]-0x30 < 0 || input[i]-0x30 > 9 || input[j]-0x30 < 0 || input[j]-0x30 > 9 {
			continue
		}
		if input[i] == input[j] {
			sum += int(input[i] - 0x30)
		}
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
