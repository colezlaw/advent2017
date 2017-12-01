package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func findSum(input []byte) int {
	sum := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i]-0x30 < 0 || input[i]-0x30 > 9 || input[i+1]-0x30 < 0 || input[i+1]-0x30 > 9 {
			continue
		}
		if input[i] == input[i+1] {
			sum += int(input[i] - 0x30)
		}
	}
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
