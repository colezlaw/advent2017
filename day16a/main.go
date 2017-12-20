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

func spin(s string, n int) string {
	r := s[len(s)-n:] + s[0:len(s)-n]
	return r
}

func partner(s string, a, b byte) string {
	var n, o int
	for i := 0; i < len(s); i++ {
		if s[i] == a {
			n = i
		}
		if s[i] == b {
			o = i
		}
	}
	return exchange(s, n, o)
}

func exchange(s string, n, o int) string {
	r := []byte(s)
	r[o], r[n] = r[n], r[o]
	return string(r)
}

func operate(s, op string) string {
	switch op[0] {
	case 's':
		n, _ := strconv.Atoi(op[1:])
		return spin(s, n)
	case 'x':
		pieces := strings.Split(op[1:], "/")
		n, _ := strconv.Atoi(pieces[0])
		o, _ := strconv.Atoi(pieces[1])
		return exchange(s, n, o)
	case 'p':
		return partner(s, op[1], op[3])
	}

	return s
}

func main() {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	start := "abcdefghijklmnop"

	s := scanner{input: string(f)}
	for {
		tok, err := s.token()
		if err != nil {
			break
		}
		start = operate(start, tok)
	}

	fmt.Println(start)
}
