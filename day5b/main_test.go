package main

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	tt := []struct {
		program  []int
		pc       int
		expected int
		err      error
	}{
		{[]int{0, 3, 0, 1, -3}, 0, 0, nil},
		{[]int{1, 3, 0, 1, -3}, 0, 1, nil},
		{[]int{2, 3, 0, 1, -3}, 1, 4, nil},
		{[]int{2, 2, 0, 1, -3}, 4, 1, nil},
		{[]int{2, 2, 0, 1, -2}, 1, 3, nil},
		{[]int{2, 3, 0, 1, -2}, 3, 4, nil},
		{[]int{2, 3, 0, 2, -2}, 4, 2, nil},
		{[]int{2, 3, 0, 2, -1}, 2, 2, nil},
		{[]int{2, 3, 1, 2, -1}, 2, 3, nil},
		{[]int{2, 3, 3, 2, -1}, 3, 0, EOP},
	}

	for x, tc := range tt {
		t.Run(fmt.Sprintf("Instruction %d", x), func(t *testing.T) {
			oldp := make([]int, len(tc.program))
			copy(oldp, tc.program)
			pc, err := step(tc.program, tc.pc)
			if err != nil {
				if err == tc.err {
					return
				}
				t.Fatalf("Unexpected error %v", err)
			}
			if tc.expected != pc {
				t.Errorf("Expected new PC to be %d, got %d", tc.expected, pc)
			}
		})
	}
}

func TestProgram(t *testing.T) {
	// Given a test program, just count the number of steps
	// it takes to jump outside the program
	tt := []struct {
		program  []int
		expected int
	}{
		{[]int{0, 3, 0, 1, -3}, 10},
	}

	for x, tc := range tt {
		t.Run(fmt.Sprintf("Program %d", x), func(t *testing.T) {
			actual := count(tc.program)
			if actual != tc.expected {
				t.Errorf("Expected to jump in %d steps, got %d", tc.expected, actual)
			}
		})
	}
}
