package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func isValid(input string) bool {
	words := make(map[string]bool)
	for _, word := range strings.Fields(strings.Trim(input, " \n")) {
		if _, ok := words[word]; ok {
			return false
		}
		words[word] = true
	}
	return true
}

func countValid(r io.Reader) int {
	count := 0
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		if isValid(s.Text()) {
			count++
		}
	}

	return count
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("opening file: %v", err)
	}
	defer f.Close()

	count := countValid(f)
	fmt.Println(count)
}
