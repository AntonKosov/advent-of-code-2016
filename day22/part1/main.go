package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2016/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []node {
	lines := aoc.ReadAllInput()
	lines = lines[2 : len(lines)-1]

	nodes := make([]node, len(lines))
	for i, line := range lines {
		ints := aoc.ParseInts(line)
		nodes[i] = node{used: ints[3], available: ints[4]}
	}

	return nodes
}

func process(nodes []node) int {
	count := 0
	for i, n1 := range nodes {
		if n1.used == 0 {
			continue
		}
		for j, n2 := range nodes {
			if i != j && n1.used <= n2.available {
				count++
			}
		}
	}
	return count
}

type node struct {
	used      int
	available int
}
