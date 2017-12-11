package main

import "testing"

func TestSamples(t *testing.T) {
	tt := []struct {
		input    string
		expected string
	}{
		{"", "a2582a3a0e66e6e86e3812dcb672a272"},
		{"AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd"},
		{"1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d"},
		{"1,2,4", "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			c := newChain(256, []byte(tc.input))
			actual := c.hash()
			if actual != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, actual)
			}
		})
	}
}
