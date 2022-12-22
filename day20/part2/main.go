package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/AntonKosov/advent-of-code-2016/aoc"
)

const (
	minValue = 0
	maxValue = 4294967295
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() []ipRange {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	ipRanges := make([]ipRange, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, "-")
		ipRanges[i] = ipRange{from: aoc.StrToInt(parts[0]), to: aoc.StrToInt(parts[1])}
	}

	return ipRanges
}

func process(ipRanges []ipRange) int {
	sort.Slice(ipRanges, func(i, j int) bool { return ipRanges[i].from < ipRanges[j].from })
	prevMax := minValue - 1
	allowed := 0
	for _, r := range ipRanges {
		allowed += aoc.Max(0, r.from-prevMax-1)
		prevMax = aoc.Max(prevMax, r.to)
	}

	return allowed
}

type ipRange struct {
	from, to int
}
