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
