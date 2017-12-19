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

// Layer defines a layer of the firewall
type Layer struct {
	Depth   int
	Range   int
	Dir     int
	Scanner int
}

func (l *Layer) String() string {
	return fmt.Sprintf("Layer{Depth: %d, Range: %d, Scanner: %d", l.Depth, l.Range, l.Scanner)
}

// Firewall is a collection of layers
type Firewall map[int]*Layer

// Advance advances all the layers in the firewall
func (f *Firewall) Advance() {
	for _, v := range map[int]*Layer(*f) {
		v.advance()
	}
}

// NewLayer does as advertised
func NewLayer(depth, rnge int) *Layer {
	return &Layer{Depth: depth, Range: rnge, Dir: 1}
}

func (l *Layer) advance() {
	l.Scanner += l.Dir
	if l.Scanner == 0 {
		l.Dir = 1
	}
	if l.Scanner == l.Range-1 {
		l.Dir = -1
	}
}

func parseLine(s string) *Layer {
	p := strings.Split(s, ": ")
	depth, _ := strconv.Atoi(p[0])
	rnge, _ := strconv.Atoi(p[1])

	l := NewLayer(depth, rnge)
	return l
}

func parse(r io.Reader) (max int, firewall Firewall) {
	firewall = make(map[int]*Layer)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		r := parseLine(s.Text())
		if r.Depth > max {
			max = r.Depth
		}
		firewall[r.Depth] = r
	}

	return
}

// Traverse sends a packet across range 0 of each layer
func (f *Firewall) Traverse(depth int) int {
	score := 0
	for i := 0; i <= depth; i++ {
		if l, ok := map[int]*Layer(*f)[i]; ok {
			fmt.Printf("Layer %d, Scanner %d\n", i, l.Scanner)
			if l.Scanner == 0 {
				fmt.Printf("BAM! %v\n", l)
				score += l.Depth * l.Range
			}
		} else {
			fmt.Printf("No rule at layer %d\n", i)
		}
		f.Advance()
	}

	return score
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening %v", err)
	}

	max, fw := parse(f)
	score := fw.Traverse(max)
	fmt.Printf("Score was %d\n", score)
}
