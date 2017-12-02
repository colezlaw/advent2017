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

func rowsum(input []int) int {
	for ix, i := range input {
		for jx, j := range input {
			if ix == jx {
				continue
			}
			if i%j == 0 {
				return i / j
			}
			if j%i == 0 {
				return j / i
			}
		}
	}
	return 0
}

func checksum(input [][]int) int {
	checksum := 0
	for _, row := range input {
		checksum += rowsum(row)
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
