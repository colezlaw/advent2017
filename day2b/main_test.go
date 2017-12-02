package main

import (
	"strings"
	"testing"
)

func TestDiff(t *testing.T) {
	tt := []struct {
		name  string
		input []int
		diff  int
	}{
		{"row1", []int{5, 9, 2, 8}, 4},
		{"row2", []int{9, 4, 7, 3}, 3},
		{"row3", []int{3, 8, 6, 5}, 2},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			diff := rowsum(tc.input)
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
			{5, 9, 2, 8},
			{9, 4, 7, 3},
			{3, 8, 6, 5},
		}, 9},
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
	input := `5	9	2	8
9	4	7	3
3	8	6	5`
	expected := 9
	result, err := rChecksum(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	if expected != result {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
