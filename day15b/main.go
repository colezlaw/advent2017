package main

import "fmt"

type rand struct {
	seed uint64
	fact uint64
	div  uint64
}

func (r *rand) rawNext() uint64 {
	r.seed = (r.seed * r.fact) % 2147483647
	return r.seed
}

func (r *rand) next() uint16 {
	for r.rawNext(); r.seed%r.div != 0; r.rawNext() {
	}
	return uint16(r.seed & 0xffff)
}

func main() {
	genA := rand{seed: 679, fact: 16807, div: 4}
	genB := rand{seed: 771, fact: 48271, div: 8}
	count := 0
	for i := 0; i < 5*1000*1000; i++ {
		if genA.next() == genB.next() {
			count++
		}
	}
	fmt.Printf("Final match count %d\n", count)
}
