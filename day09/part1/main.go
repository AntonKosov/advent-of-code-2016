package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2016/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() []byte {
	lines := aoc.ReadAllInput()
	return []byte(lines[0])
}

func process(data []byte) int {
	length := 0
	for i := 0; i < len(data); i++ {
		v := data[i]
		if v != '(' {
			length++
			continue
		}
		seqLen, repeat, closeBracket := readParams(data, i)
		length += seqLen * repeat
		i = closeBracket + seqLen
	}

	return length
}

func readParams(data []byte, start int) (length, repeat, closeBracket int) {
	i := start + 1
	for ; data[i] != 'x'; i++ {
		length = length*10 + int(data[i]-'0')
	}

	i++
	for ; data[i] != ')'; i++ {
		repeat = repeat*10 + int(data[i]-'0')
	}

	return length, repeat, i
}
