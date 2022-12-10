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

func read() [][]direction {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	instructions := make([][]direction, len(lines))
	for i, line := range lines {
		dirs := make([]direction, len(line))
		for j, d := range line {
			dirs[j] = direction(d)
		}
		instructions[i] = dirs
	}

	return instructions
}

func process(instructions [][]direction) int {
	const size = 3
	x, y := 1, 1
	code := 0
	for _, dirs := range instructions {
		for _, dir := range dirs {
			switch dir {
			case upDirection:
				if y > 0 {
					y--
				}
			case downDirection:
				if y < size-1 {
					y++
				}
			case leftDirection:
				if x > 0 {
					x--
				}
			case rightDirection:
				if x < size-1 {
					x++
				}
			default:
				panic(fmt.Sprintf("unknown direction: %v", dir))
			}
		}
		code = code*10 + y*size + x + 1
	}

	return code
}

type direction rune

const (
	upDirection    direction = 'U'
	downDirection  direction = 'D'
	rightDirection direction = 'R'
	leftDirection  direction = 'L'
)
