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

func parseLine(s string) []int {
	parts := strings.Split(s, " <-> ")
	r := strings.Split(parts[1], ", ")

	result := make([]int, len(r)+1)
	result[0], _ = strconv.Atoi(parts[0])
	for i, rs := range r {
		result[i+1], _ = strconv.Atoi(rs)
	}

	return result
}

// SECURITY PEOPLE BEWARE. TIMING ATTACK
func contains(haystack []int, needle int) bool {
	for _, n := range haystack {
		if n == needle {
			return true
		}
	}
	return false
}

// SECURITY PEOPLE DOUBLY AWARE
func containsAny(haystack []int, needle []int) bool {
	for _, n := range needle {
		if contains(haystack, n) {
			return true
		}
	}
	return false
}

func ohgroup(r io.Reader) []int {
	// Unique list of numbers we've seen
	result := map[int]bool{0: true}

	allseries := make([][]int, 0)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		allseries = append(allseries, parseLine(s.Text()))
	}

	// While we continue to find new items
	newlen, prevlen := 1, 0
	for newlen > prevlen {
		prevlen = len(result)

		for _, r := range allseries {
		row:
			for _, c := range r {
				if _, ok := result[c]; ok {
					for _, n := range r {
						result[n] = true
					}
					break row
				}
			}
		}

		newlen = len(result)
	}

	list := make([]int, len(result))
	i := 0
	for k := range result {
		list[i] = k
		i++
	}

	return list
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening: %v", err)
	}
	defer f.Close()

	result := ohgroup(f)
	fmt.Printf("%v\n", result)
	fmt.Printf("len %d\n", len(result))
}
