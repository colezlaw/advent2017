package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

type scanner struct {
	input string
	start int
	pos   int
}

func (s *scanner) token() (string, error) {
	if s.pos >= len(s.input) {
		return "", io.EOF
	}
	for ; s.pos < len(s.input); s.pos++ {
		if s.input[s.pos] == ',' {
			ret := s.input[s.start:s.pos]
			s.pos = s.pos + 1
			s.start = s.pos
			return ret, nil
		}
	}
	if s.start < s.pos {
		return s.input[s.start:s.pos], nil
	}
	return "", io.EOF
}

func spin(s []byte, l int) {
	result := append(s[len(s)-l:], s[0:len(s)-l]...)
	copy(s, result)
}

func partner(s []byte, a, b byte) {
	var n, o int
	for i := 0; i < len(s); i++ {
		if s[i] == a {
			n = i
		}
		if s[i] == b {
			o = i
		}
	}
	exchange(s, n, o)
}

func exchange(s []byte, n, o int) {
	s[o], s[n] = s[n], s[o]
}

func operate(s []byte, op string) {
	switch op[0] {
	case 's':
		n, _ := strconv.Atoi(op[1:])
		spin(s, n)
	case 'x':
		pieces := strings.Split(op[1:], "/")
		n, _ := strconv.Atoi(pieces[0])
		o, _ := strconv.Atoi(pieces[1])
		exchange(s, n, o)
	case 'p':
		partner(s, op[1], op[3])
	}
}

func main() {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	start := []byte("abcdefghijklmnop")

	// See if we can find a cycle
	i := 1
	for i = 1; ; i++ {
		s := scanner{input: string(f)}
		for {
			tok, err := s.token()
			if err != nil {
				break
			}
			operate(start, tok)
		}
		if string(start) == "abcdefghijklmnop" {
			break
		}
	}

	// We found a cycle in this many steps. So
	// we can short-circuit the race to a billion
	// by taking however many dances it would take
	// AFTER the next nearest number of cycles, and
	// just running that many times.
	times := (1 * 1000 * 1000 * 1000) % i
	for j := 0; j < times; j++ {
		s := scanner{input: string(f)}
		for {
			tok, err := s.token()
			if err != nil {
				break
			}
			operate(start, tok)
		}
	}
	fmt.Printf("%s\n", start)
}
