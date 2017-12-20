package main

import "testing"

func TestOne(t *testing.T) {
	genA := rand{seed: 65, fact: 16807, div: 4}
	expected := []uint16{
		uint16(1352636452 & 0xffff),
		uint16(1992081072 & 0xffff),
		uint16(530830436 & 0xffff),
		uint16(1980017072 & 0xffff),
		uint16(740335192 & 0xffff),
	}

	for _, tc := range expected {
		r := genA.next()
		if r != tc {
			t.Errorf("Expected %d, got %d (%d)", tc, genA.seed, r)
		}
	}
}

func TestCompare(t *testing.T) {
	genA := rand{div: 4, fact: 16807, seed: 65}
	genB := rand{div: 8, fact: 48271, seed: 8921}
	counter := 0
	for {
		counter++
		if genA.next() == genB.next() {
			break
		}
	}
	if counter != 1056 {
		t.Errorf("Expected match at 1056 iterations, got %d", counter)
	}
}
