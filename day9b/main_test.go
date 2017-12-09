package main

import "testing"

func TestBraces(t *testing.T) {
	tt := []struct {
		input    string
		expected int
	}{
		{"{}", 1},
		{"{{}}", 3},
		{"{{},{},{}}", 7},
		{"{{},{},<{}>}", 5},
		{"{{{}}}", 6},
		{"{{},{}}", 5},
		{"{{{},{},{{}}}}", 16},
		{"{<a>,<a>,<a>,<a>}", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			l := newLexer(tc.input)
			actual := l.lex()
			if actual != tc.expected {
				t.Errorf("Expected score of %d, got %d", tc.expected, actual)
			}
		})
	}
}

func TestGarbage(t *testing.T) {
	tt := []struct {
		input    string
		expected int
	}{
		{"{<>}", 0},
		{"{<random characters>}", 17},
		{"{<<<<>}", 3},
		{"{<{!>}>}", 2},
		{"{<!!>}", 0},
		{"{<!!!>>}", 0},
		{"{<{o\"o!a,<{i<a>}", 10},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			l := newLexer(tc.input)
			n := l.lex()
			if n != 1 {
				t.Errorf("Expected 1 count, got %d", n)
			}
			if l.garbageCount != tc.expected {
				t.Errorf("Expected garbage count of %d, got %d", tc.expected, l.garbageCount)
			}
		})
	}
}
