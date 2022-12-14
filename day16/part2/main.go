package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2016/aoc"
)

func main() {
	state, size := read()
	r := process(state, size)
	fmt.Printf("Answer: %v\n", r)
}

func read() (state []byte, size int) {
	lines := aoc.ReadAllInput()

	st := []byte(lines[0])
	for _, v := range st {
		state = append(state, v-'0')
	}

	return state, aoc.StrToInt(lines[1])
}

func process(state []byte, size int) string {
	generateCurve(&state, size)
	state = state[:size]
	for len(state)&1 == 0 {
		for i := 0; i < len(state)/2; i++ {
			v1 := state[i*2]
			v2 := state[i*2+1]
			h := 0
			if (v1+v2)&1 == 0 {
				h = 1
			}
			state[i] = byte(h)
		}
		state = state[:len(state)/2]
	}

	var checksum strings.Builder
	for _, s := range state {
		checksum.WriteByte(s + '0')
	}

	return checksum.String()
}

func generateCurve(state *[]byte, size int) {
	for len(*state) < size {
		c := make([]byte, len(*state))
		for i, v := range *state {
			c[len(*state)-i-1] = 1 - v
		}
		*state = append(*state, 0)
		*state = append(*state, c...)
	}
}
