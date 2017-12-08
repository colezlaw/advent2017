package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func max(m map[string]int) int {
	ret := int(math.MinInt32)
	for _, v := range m {
		if v > ret {
			ret = v
		}
	}

	return ret
}

func compute(input string, registers map[string]int) (tmax int) {
	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(bufio.ScanLines)

	for s.Scan() {
		fields := strings.Fields(s.Text())

		// It's conceivable that we could check against a register (which would return 0)
		// but not explicitly set it - in which case maybe max could be negative, but it
		// should be 0
		if _, ok := registers[fields[4]]; !ok {
			registers[fields[4]] = 0
		}

		incr := false
		test, _ := strconv.Atoi(fields[6])
		switch fields[5] {
		case "<":
			incr = registers[fields[4]] < test
		case "<=":
			incr = registers[fields[4]] <= test
		case ">":
			incr = registers[fields[4]] > test
		case ">=":
			incr = registers[fields[4]] >= test
		case "!=":
			incr = registers[fields[4]] != test
		case "==":
			incr = registers[fields[4]] == test
		}
		mul := 1
		if fields[1] == "dec" {
			mul = -1
		}

		if incr {
			value, _ := strconv.Atoi(fields[2])
			registers[fields[0]] += mul * value
			if registers[fields[0]] > tmax {
				tmax = registers[fields[0]]
			}
		}
	}

	return
}

func main() {
	buff, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	result := make(map[string]int)
	tmax := compute(string(buff), result)
	fmt.Println(max(result), tmax)
}
