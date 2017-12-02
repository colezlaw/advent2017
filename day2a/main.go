package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func minmax(input []int) (min, max int) {
	if input == nil || len(input) == 0 {
		return
	}
	min, max = input[0], input[0]
	for _, n := range input {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	return
}

func rowdiff(input []int) int {
	min, max := minmax(input)
	return max - min
}

func checksum(input [][]int) int {
	checksum := 0
	for _, row := range input {
		checksum += rowdiff(row)
	}
	return checksum
}

func splitLine(input string) ([]int, error) {
	var result []int
	for _, s := range strings.Split(strings.Trim(input, " \n"), "\t") {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}

	return result, nil
}

func read(r io.Reader) ([][]int, error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	result := make([][]int, 0)
	for s.Scan() {
		line := s.Text()
		ints, err := splitLine(line)
		if err != nil {
			return nil, err
		}
		result = append(result, ints)
	}
	if s.Err() != nil {
		return nil, s.Err()
	}
	return result, nil
}

func rChecksum(r io.Reader) (int, error) {
	vec, err := read(r)
	if err != nil {
		return 0, err
	}
	return checksum(vec), nil
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error opening %v", err)
	}
	defer f.Close()

	fmt.Println(rChecksum(f))
}
