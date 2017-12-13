package main

import "testing"

func TestParseLine(t *testing.T) {
	tt := []struct {
		line     string
		expected []int
	}{
		{"0 <-> 2", []int{0, 2}},
		{"1 <-> 1", []int{1, 1}},
		{"2 <-> 0, 3, 4", []int{2, 0, 3, 4}},
		{"3 <-> 2, 4", []int{3, 2, 3}},
		{"4 <-> 2, 3, 6", []int{4, 2, 3, 6}},
		{"5 <-> 6", []int{5, 6}},
		{"6 <-> 4, 5", []int{6, 4, 5}},
	}

	for _, tc := range tt {
		t.Run(tc.line, func(t *testing.T) {
			actual := parseLine(tc.line)
			if len(actual) != len(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
