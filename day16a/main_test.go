package main

import (
	"testing"
)

func TestSpin(t *testing.T) {
	start := "abcde"
	expected := "cdeab"
	actual := spin(start, 3)
	if actual != expected {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}
}

func TestExchange(t *testing.T) {
	start := "eabcd"
	expected := "eabdc"
	actual := exchange(start, 3, 4)
	if actual != expected {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}
}

func TestPartner(t *testing.T) {
	start := "eabdc"
	expected := "baedc"
	actual := partner(start, 'b', 'e')
	if actual != expected {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}
}

func TestOperate(t *testing.T) {
	input := "abcde"
	tt := []struct {
		op       string
		expected string
	}{
		{"s1", "eabcd"},
		{"s3", "cdeab"},
		{"x2/4", "abedc"},
		{"pa/b", "bacde"},
	}

	for _, tc := range tt {
		t.Run(tc.op, func(t *testing.T) {
			r := operate(input, tc.op)
			if r != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, r)
			}
		})
	}
}

func TestScanner(t *testing.T) {
	s := scanner{input: "s1,x3/4,pb/e"}
	expected := 3
	count := 0
	for {
		_, err := s.token()
		if err != nil {
			break
		}
		count++
	}
	if count != expected {
		t.Errorf("Expected %d, got %d", expected, count)
	}
}
