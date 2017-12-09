package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type state uint8

const (
	stateBrace state = iota
	stateGarbage
	stateBang
	stateStart
)

type lexer struct {
	state        // State we're in so we know how to handle the next character
	input string // input string
	pos   int    // where we are in the input
	score int    // accumulated score
	depth int    // how many levels nested we are in groups
}

func newLexer(input string) *lexer {
	l := &lexer{
		state: stateStart,
		input: input,
	}
	return l
}

func (l *lexer) lex() int {
	for ; l.pos < len(l.input); l.pos++ {
		// Check the current character
		switch l.input[l.pos] {
		case '{':
			if l.state == stateStart || l.state == stateBrace {
				// Increment the depth
				l.depth++
				l.state = stateBrace
			}
		case '}':
			if l.state == stateBrace {
				l.score += l.depth
				l.depth--
			}
		case '<':
			if l.state == stateBrace {
				l.consumeGarbage()
			}
		case '!':
			// Consume one more byte, no matter what's after here
			l.pos++
		}
	}

	return l.score
}

// Garbage needs to be handled somewhat specially
func (l *lexer) consumeGarbage() {
	if l.input[l.pos] != '<' {
		return
	}
	l.pos++
	for ; l.pos <= len(l.input); l.pos++ {
		switch l.input[l.pos] {
		case '!':
			// Consume one more byte
			l.pos++
		case '>':
			// Current garbage is over, return
			return
		}
	}
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("error reading %v", err)
	}

	l := newLexer(string(input))
	fmt.Println(l.lex())
}
