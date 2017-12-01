package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func findSum(input []byte) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		next := (i + len(input)/2) % len(input)
		if input[i]-0x30 < 0 || input[i]-0x30 > 9 || input[next]-0x30 < 0 || input[next]-0x30 > 9 {
			continue
		}
		if input[i] == input[next] {
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
