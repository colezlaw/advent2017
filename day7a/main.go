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

type node struct {
	ident    string
	weight   int
	children []string
}

func parseNode(input string) node {
	fields := strings.Fields(input)
	n := node{}
	n.ident = fields[0]
	n.weight, _ = strconv.Atoi(fields[1][1 : len(fields[1])-1])
	if len(fields) > 3 {
		n.children = make([]string, len(fields)-3)
		for i := 3; i < len(fields)-1; i++ {
			n.children[i-3] = fields[i][:len(fields[i])-1]
		}
		n.children[cap(n.children)-1] = fields[len(fields)-1]
	}

	return n
}

func parseReader(r io.Reader) []node {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	var ret []node

	for s.Scan() {
		ret = append(ret, parseNode(s.Text()))
	}

	return ret
}

func getRoots(nodes []node) []string {
	descendants := make(map[string]bool)
	for _, n := range nodes {
		// The ident is not a child
		if _, ok := descendants[n.ident]; !ok {
			descendants[n.ident] = false
		}

		// But all the children are children
		for _, s := range n.children {
			descendants[s] = true
		}
	}

	// Now build a list of all the seen idents that were not children
	ret := make([]string, 0)
	for k, v := range descendants {
		if !v {
			ret = append(ret, k)
		}
	}

	return ret
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error opening %v", err)
	}
	defer f.Close()

	nodes := parseReader(f)
	fmt.Println(getRoots(nodes))
}
