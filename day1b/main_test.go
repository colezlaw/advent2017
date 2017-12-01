package main

import "testing"

func TestDay1(t *testing.T) {
	testCases := []struct {
		input    []byte
		expected int
	}{
		{[]byte("1212"), 6},
		{[]byte("1221"), 0},
		{[]byte("123425"), 4},
		{[]byte("123123"), 12},
		{[]byte("12131415"), 4},
	}
	for _, tc := range testCases {
		t.Run(string(tc.input), func(t *testing.T) {
			if findSum(tc.input) != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, findSum(tc.input))
			}
		})
	}
}
