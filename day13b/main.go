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

func parseLine(s string) (int, int) {
	p := strings.Split(s, ": ")
	depth, _ := strconv.Atoi(p[0])
	rnge, _ := strconv.Atoi(p[1])

	// Here's the real magic for this one. Since we don't
	// need to know what the score is, only when we find a route
	// all the way through, we only care whether a layer is
	// scanning on its depth 0.
	rnge = rnge*2 - 2

	return depth, rnge
}

func parse(r io.Reader) (max int, firewall map[int]int) {
	firewall = make(map[int]int)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		depth, rnge := parseLine(s.Text())
		if depth > max {
			max = depth
		}
		firewall[depth] = rnge
	}

	return
}

func traverse(max int, fw map[int]int) int {
loop:
	for delay := 0; ; delay++ { // Loop to delay
		// fmt.Printf("Delay of %d\n", delay)
		for layer := 0; layer <= max; layer++ {
			// fmt.Printf("  Layer %d\n", layer)
			if v, ok := fw[layer]; ok { // Rule on this layer. v is the ful;
				// fmt.Printf("    Depth of layer %d\n", v)
				// length of a scan.
				if (layer+delay)%v == 0 {
					continue loop
				}
			}
		}
		// Made it all the way through
		return delay
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening %v", err)
	}

	max, fw := parse(f)

	count := traverse(max, fw)
	fmt.Println(count)
}
