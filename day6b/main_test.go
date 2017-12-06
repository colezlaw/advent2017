package main

import (
	"testing"
)

func TestOne(t *testing.T) {
	tt := []struct {
		name     string
		banks    []int
		expected int
	}{
		{"simplemax", []int{3, 4, 5}, 2},
		{"example", []int{0, 2, 7, 0}, 2},
		{"duplicate", []int{7, 2, 7, 0}, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := maxindex(tc.banks)
			if actual != tc.expected {
				t.Errorf("Expected index %d, got %d", tc.expected, actual)
			}
		})
	}
}

func TestDistribute(t *testing.T) {
	tt := []struct {
		name     string
		banks    []int
		expected []int
	}{
		{"simplemax", []int{3, 4, 5}, []int{5, 6, 1}},
		{"example", []int{0, 2, 7, 0}, []int{2, 4, 1, 2}},
		{"duplicate", []int{7, 2, 7, 0}, []int{1, 4, 9, 2}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Merely distribute one time
			cp := make([]int, len(tc.expected))
			copy(cp, tc.banks)
			distribute(cp)
			for k, v := range cp {
				if v != tc.expected[k] {
					t.Errorf("Expected index %d to be %d, got %d", k, tc.expected[k], v)
				}
			}
		})
	}
}
func TestOptimize(t *testing.T) {
	// This tests optimize, which optimizes the distribution but stops when it
	// sees a duplicate arrangement
	tt := []struct {
		name     string
		banks    []int
		expected []int
		rounds   int
	}{
		{"example", []int{0, 2, 7, 0}, []int{2, 4, 1, 2}, 5},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			cp := make([]int, len(tc.banks))
			copy(cp, tc.banks)
			rounds := optimize(cp)
			if rounds != tc.rounds {
				t.Errorf("Expected to go %d rounds, got %d", tc.rounds, rounds)
			}
			for k, v := range cp {
				if tc.expected[k] != v {
					t.Errorf("Expected index %d to be %d, got %d", k, tc.expected[k], v)
				}
			}
		})
	}
}
