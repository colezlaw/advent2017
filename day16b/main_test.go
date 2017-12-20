package main

import (
	"testing"
)

func TestSpin(t *testing.T) {
	start := []byte("abcde")
	expected := "cdeab"
	spin(start, 3)
	if string(start) != expected {
		t.Errorf("Expected '%s', got '%s'", expected, start)
	}
}

func TestExchange(t *testing.T) {
	start := []byte("eabcd")
	expected := "eabdc"
	exchange(start, 3, 4)
	if string(start) != expected {
		t.Errorf("Expected '%s', got '%s'", expected, start)
	}
}

func TestPartner(t *testing.T) {
	start := []byte("eabdc")
	expected := "baedc"
	partner(start, 'b', 'e')
	if string(start) != expected {
		t.Errorf("Expected '%s', got '%s'", expected, start)
	}
}

func TestOperate(t *testing.T) {
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
		input := []byte("abcde")
		t.Run(tc.op, func(t *testing.T) {
			operate(input, tc.op)
			if string(input) != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, input)
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
