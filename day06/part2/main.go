package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2016/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() [][]byte {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	message := make([][]byte, len(lines))
	for i, line := range lines {
		message[i] = []byte(line)
	}

	return message
}

func process(data [][]byte) string {
	n := len(data[0])
	var word strings.Builder
	for i := 0; i < n; i++ {
		freq := map[byte]int{}
		for _, l := range data {
			freq[l[i]]++
		}

		minFreq := len(data) + 1
		var ch byte
		for r, c := range freq {
			if minFreq > c {
				minFreq = c
				ch = r
			}
		}

		word.WriteByte(ch)
	}

	return word.String()
}
