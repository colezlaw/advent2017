package main

import (
	"strings"
	"testing"
)

func TestMinMax(t *testing.T) {
	tt := []struct {
		name  string
		input []int
		max   int
		min   int
	}{
		{"row1", []int{5, 1, 9, 5}, 9, 1},
		{"row2", []int{7, 5, 3}, 7, 3},
		{"row3", []int{2, 4, 6, 8}, 8, 2},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			min, max := minmax(tc.input)
			if min != tc.min {
				t.Errorf("Expected min to be %d, got %d", tc.min, min)
			}
			if max != tc.max {
				t.Errorf("Expected max to be %d, got %d", tc.max, max)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	tt := []struct {
		name  string
		input []int
		diff  int
	}{
		{"row1", []int{5, 1, 9, 5}, 8},
		{"row2", []int{7, 5, 3}, 4},
		{"row3", []int{2, 4, 6, 8}, 6},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			diff := rowdiff(tc.input)
			if diff != tc.diff {
				t.Errorf("Expected diff to be %d, got %d", tc.diff, diff)
			}
		})
	}
}

func TestChecksum(t *testing.T) {
	tt := []struct {
		name     string
		input    [][]int
		checksum int
	}{
		{"example", [][]int{
			{5, 1, 9, 5},
			{7, 5, 3},
			{2, 4, 6, 8},
		}, 18},
	}
	for _, tc := range tt {
		r := checksum(tc.input)
		if r != tc.checksum {
			t.Errorf("Expected checksum to be %d, got %d", tc.checksum, r)
		}
	}
}

func TestSplitLine(t *testing.T) {
	tt := []struct {
		input string
		len   int
		first int
	}{
		{"5\t1\t9\t5", 4, 5},
		{"7\t5\t3", 3, 7},
		{"2\t4\t6\t8", 4, 2},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			result, err := splitLine(tc.input)
			if err != nil {
				t.Fatalf("Unexpected error %v", err)
			}
			if len(result) != tc.len {
				t.Errorf("Expected len to be %d, got %d", tc.len, len(result))
			}
			if result[0] != tc.first {
				t.Errorf("Expected first to be %d, got %d", tc.first, result[0])
			}
		})
	}
}

func TestRead(t *testing.T) {
	input := `5	1	9	5
7	5	3
2	4	6	8`
	expected := [][]int{
		{5, 1, 9, 5},
		{7, 5, 3},
		{2, 4, 6, 8},
	}

	result, err := read(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	if len(result) != len(expected) {
		t.Fatalf("Expected len to be %d, got %d", len(expected), len(result))
	}
	for ri, row := range result {
		if len(row) != len(expected[ri]) {
			t.Fatalf("Expected len of row %d to be %d, got %d", ri, len(expected[ri]), len(row))
		}
	}
}

func TestRChecksum(t *testing.T) {
	input := `5	1	9	5
7	5	3
2	4	6	8`
	expected := 18
	result, err := rChecksum(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	if expected != result {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
