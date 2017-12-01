package main

import "testing"

func TestDay1(t *testing.T) {
	testCases := []struct {
		input    []byte
		expected int
	}{
		{[]byte("1122"), 3},
		{[]byte("1111"), 4},
		{[]byte("1234"), 0},
		{[]byte("91212129"), 9},
	}
	for _, tc := range testCases {
		t.Run(string(tc.input), func(t *testing.T) {
			if findSum(tc.input) != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, findSum(tc.input))
			}
		})
	}
}
