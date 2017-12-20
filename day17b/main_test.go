package main

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	sl := newSpinlock(382)
	for i := 0; i < 50*1000*1000; i++ {
		sl.spin()
	}
	fmt.Println(sl.slot1)
}
