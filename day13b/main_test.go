package main

import (
	"strings"
	"testing"
)

func TestStrategic(t *testing.T) {
	f := Firewall{
		0: NewLayer(0, 3),
		1: NewLayer(1, 2),
		4: NewLayer(4, 4),
		6: NewLayer(6, 4),
	}

	for i := 0; i < 10; i++ {
		f.Advance()
	}
	score := f.Traverse(6)
	if score != 0 {
		t.Errorf("Expected score 0, got %d", score)
	}
}

func TestOne(t *testing.T) {
	tt := []struct {
		input    string
		expected Layer
	}{
		{"0: 3", Layer{Depth: 0, Range: 3}},
		{"1: 2", Layer{Depth: 1, Range: 2}},
		{"4: 4", Layer{Depth: 4, Range: 4}},
		{"6: 4", Layer{Depth: 6, Range: 4}},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			result := parseLine(tc.input)
			if result.Depth != tc.expected.Depth {
				t.Errorf("Expected depth %d, got %d", tc.expected.Depth, result.Depth)
			}
		})
	}
}

func TestParse(t *testing.T) {
	input := `0: 3
1: 2
4: 4
6: 4`
	expected := Firewall{
		0: NewLayer(0, 3),
		1: NewLayer(1, 2),
		4: NewLayer(4, 4),
		6: NewLayer(6, 4),
	}

	max, fw := parse(strings.NewReader(input))
	if max != 6 {
		t.Fatalf("Expected max to be 6, got %d", max)
	}

	if len(fw) != len(expected) {
		t.Fatalf("Expected len to be %d, got %d", len(expected), len(fw))
	}

	for k, v := range expected {
		if gk, ok := fw[k]; !ok {
			t.Fatalf("Expected %d to be set, was not", k)
		} else {
			if gk.String() != v.String() {
				t.Errorf("Expected %d to be %s, got %s", k, v, gk)
			}
		}
	}
}

func TestScan(t *testing.T) {
	l := NewLayer(0, 3)
	l.Scanner = 1
	l.advance() // Should be 2
	if l.Scanner != 2 {
		t.Fatalf("Expected 2, got %d", l.Scanner)
	}
	l.advance() // Should be 1
	if l.Scanner != 1 {
		t.Fatalf("Expected 1, got %d", l.Scanner)
	}
	l.advance() // Should be 0
	if l.Scanner != 0 {
		t.Fatalf("Expected 0, got %d", l.Scanner)
	}
}

func TestAdvance(t *testing.T) {
	f := Firewall{
		0: NewLayer(0, 3),
		1: NewLayer(0, 2),
		4: NewLayer(0, 4),
		6: NewLayer(0, 4),
	}

	f.Advance()
	f.Advance()
	f.Advance()
	if f[0].Scanner != 1 {
		t.Errorf("Expected range 1, got %d", f[0].Scanner)
	}
	if f[1].Scanner != 1 {
		t.Errorf("Expected range 1, got %d", f[1].Scanner)
	}
	if f[4].Scanner != 3 {
		t.Errorf("Expected range 3, got %d", f[4].Scanner)
	}
}

func TestTraverse(t *testing.T) {
	f := Firewall{
		0: NewLayer(0, 3),
		1: NewLayer(1, 2),
		4: NewLayer(4, 4),
		6: NewLayer(6, 4),
	}
	score := f.Traverse(6)
	if score != 24 {
		t.Errorf("Expected 24, got %d", score)
	}
}
