package main

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	sl := newSpinlock(3)
	for i := 0; i < 10; i++ {
		sl.spin()
	}

	for i := 0; i < len(sl.cbuf); i++ {
		if sl.cbuf[i] == 2017 {
			fmt.Println(sl.cbuf[i+1])
		}
	}
}
