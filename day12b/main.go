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

func inagroup(haystack [][]int, needle int) bool {
	for _, r := range haystack {
		for _, c := range r {
			if c == needle {
				return true
			}
		}
	}

	return false
}

func boot(r io.Reader, n int) [][]int {
	allseries := make([][]int, 0)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		allseries = append(allseries, parseLine(s.Text()))
	}

	var allgroups [][]int

	// First, make the group containing 0
	allgroups = append(allgroups, agroup(allseries, 0))

	// Find the next available number to start a group
	for _, r := range allseries {
		for _, c := range r {
			if !inagroup(allgroups, c) {
				allgroups = append(allgroups, agroup(allseries, c))
			}
		}
	}

	return allgroups
}

func agroup(allseries [][]int, n int) []int {

	// Unique list of numbers we've seen
	result := map[int]bool{n: true}

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

	result := boot(f, 0)
	fmt.Println(result)
	fmt.Println(len(result))
}
