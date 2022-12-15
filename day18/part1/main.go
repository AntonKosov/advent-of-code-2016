package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2016/aoc"
)

func main() {
	traps := read()
	r := process(traps)
	fmt.Printf("Answer: %v\n", r)
}

func read() []bool {
	line := aoc.ReadAllInput()[0]

	traps := make([]bool, len(line))
	for i, v := range line {
		traps[i] = v == '^'
	}

	return traps
}

func process(traps []bool) int {
	safe := 0
	trapCases := [][3]bool{
		{true, false, false},
		{false, false, true},
		{true, true, false},
		{false, true, true},
	}

	for i := 0; i < 40; i++ {
		for _, trap := range traps {
			if !trap {
				safe++
			}
		}

		nextRow := make([]bool, len(traps))
		for j := range nextRow {
			trapCase := [3]bool{}
			from, to := aoc.Max(0, j-1), aoc.Min(len(nextRow)-1, j+1)
			for k := from; k <= to; k++ {
				trapCase[k-j+1] = traps[k]
			}

			nextRow[j] = aoc.Contains(trapCases, func(c [3]bool) bool { return c == trapCase })
		}

		traps = nextRow
	}

	return safe
}
