package main

import "testing"

func TestOne(t *testing.T) {
	genA := rand{seed: 65, fact: 16807}
	expected := []uint16{
		uint16(1092455 & 0xffff),
		uint16(1181022009 & 0xffff),
		uint16(245556042 & 0xffff),
		uint16(1744312007 & 0xffff),
		uint16(1352636452 & 0xffff),
	}

	for _, tc := range expected {
		r := genA.next()
		if r != tc {
			t.Errorf("Expected %d, got %d", tc, r)
		}
	}
}

func TestBig(t *testing.T) {
	genA := rand{seed: 65, fact: 16807}
	genB := rand{seed: 8921, fact: 48271}
	count := 0
	for i := 0; i < 40*1000*1000; i++ {
		if genA.next() == genB.next() {
			count++
		}
	}
	if count != 588 {
		t.Errorf("Expected 588, got %d", count)
	}
}
