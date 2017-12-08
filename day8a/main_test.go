package main

import "testing"

var sample = `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`

func TestSample(t *testing.T) {
	expected := map[string]int{
		"a": 1,
		"b": 0,
		"c": -10,
	}

	result := make(map[string]int)

	compute(sample, result)
	if len(expected) != len(result) {
		t.Errorf("Expected len(result) to be %d, got %d", len(expected), len(result))
	}

	for k, v := range expected {
		if v1, ok := result[k]; !ok {
			t.Errorf("Expected result[%s] to be set, but it was not", k)
		} else if v1 != v {
			t.Errorf("Expected result[%s] to be %d, got %d", k, v, v1)
		}
	}
}

func TestMax(t *testing.T) {
	expected := 1
	input := map[string]int{
		"a": 1,
		"b": 0,
		"c": -10,
	}

	actual := max(input)
	if expected != actual {
		t.Errorf("Expected max of %d, got %d", expected, actual)
	}
}
