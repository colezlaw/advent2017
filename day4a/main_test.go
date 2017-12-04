package main

import (
	"strings"
	"testing"
)

func TestOne(t *testing.T) {
	tt := []struct {
		input    string
		expected bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaa", true},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			if tc.expected != isValid(tc.input) {
				t.Errorf("Expected %t, got %t", tc.expected, isValid(tc.input))
			}
		})
	}
}

func TestCount(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected int
	}{
		{"site examples", `aa bb cc dd ee
aa bb cc dd aa
aa bb cc dd aaa`,
			2},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)
			actual := countValid(r)
			if actual != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
