package main

import "testing"

func TestReverse(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		pos      int
		l        int
		expected []int
	}{
		{"Simple", []int{1, 2, 3, 4}, 0, 4, []int{4, 3, 2, 1}},
		{"Wrap", []int{1, 2, 3, 4}, 2, 4, []int{4, 3, 2, 1}},
		{"Odd", []int{1, 2, 3, 4}, 0, 3, []int{3, 2, 1, 4}},
		{"Round1", []int{0, 1, 2, 3, 4}, 0, 3, []int{2, 1, 0, 3, 4}},
		{"Round2", []int{2, 1, 0, 3, 4}, 3, 4, []int{4, 3, 0, 1, 2}},
		{"Round3", []int{4, 3, 0, 1, 2}, 3, 1, []int{4, 3, 0, 1, 2}},
		{"Round4", []int{4, 3, 0, 1, 2}, 1, 5, []int{3, 4, 2, 1, 0}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			reverse(tc.input, tc.pos, tc.l)
			if len(tc.input) != len(tc.expected) {
				t.Fatalf("Expected len of %d, got %d", len(tc.expected), len(tc.input))
			}
			for k, v := range tc.input {
				if v != tc.expected[k] {
					t.Errorf("Expected index %d to be %d, was %d", k, tc.expected[k], v)
				}
			}
		})
	}
}

func TestRounds(t *testing.T) {
	c := newChain(5, []int{3, 4, 1, 5})
	expected := [][]int{
		{2, 1, 0, 3, 4},
		{4, 3, 0, 1, 2},
		{4, 3, 0, 1, 2},
		{3, 4, 2, 1, 0},
	}
	for i := range c.lengths {
		c.round()
		if c.skip != i+1 {
			t.Fatalf("Expected skip to be %d, was %d", i+1, c.skip)
		}
		for k, v := range expected[i] {
			if c.list[k] != v {
				t.Errorf("Expected element %d to be %d, was %d", k, v, c.list[k])
			}
		}
	}
}
