package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// EOP - End of Progran
var EOP = errors.New("EOP")

func step(program []int, pc int) (int, error) {
	// First, check to see if the offset of pc is within the slice range
	if pc < 0 || pc >= len(program) {
		return 0, EOP
	}

	// We know the current pc is in range, so follow its instruction
	newPC := pc + program[pc]
	fmt.Printf("new PC is %d\n", newPC)

	// Check to see if newPC is still within the program
	if newPC < 0 || newPC >= len(program) {
		fmt.Println("Returning EOP")
		return newPC, EOP
	}

	// Increment the previous instruction and return
	if program[pc] >= 3 {
		program[pc]--
	} else {
		program[pc]++
	}

	return newPC, nil
}

func count(program []int) int {
	count := 0
	var pc int
	var err error
	for {
		pc, err = step(program, pc)
		count++
		if err == EOP {
			break
		}
	}

	return count
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error opening: %v", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var program []int
	for s.Scan() {
		n, err := strconv.Atoi(strings.Trim(s.Text(), " \n"))
		if err != nil {
			log.Fatalf("Unexpected error parsing %v", err)
		}
		program = append(program, n)
	}

	steps := count(program)
	fmt.Println(steps)
}
