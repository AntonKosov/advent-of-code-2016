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

func process(instructions [][]direction) string {
	keypad := map[aoc.Vector2]rune{
		aoc.NewVector2(2, 0): '1',
		aoc.NewVector2(1, 1): '2',
		aoc.NewVector2(2, 1): '3',
		aoc.NewVector2(3, 1): '4',
		aoc.NewVector2(0, 2): '5',
		aoc.NewVector2(1, 2): '6',
		aoc.NewVector2(2, 2): '7',
		aoc.NewVector2(3, 2): '8',
		aoc.NewVector2(4, 2): '9',
		aoc.NewVector2(1, 3): 'A',
		aoc.NewVector2(2, 3): 'B',
		aoc.NewVector2(3, 3): 'C',
		aoc.NewVector2(2, 4): 'D',
	}
	dirs := map[direction]aoc.Vector2{
		upDirection:    aoc.NewVector2(0, -1),
		downDirection:  aoc.NewVector2(0, 1),
		leftDirection:  aoc.NewVector2(-1, 0),
		rightDirection: aoc.NewVector2(1, 0),
	}
	pos := aoc.NewVector2(0, 2)
	var code strings.Builder
	for _, directions := range instructions {
		for _, dir := range directions {
			candidate := pos.Add(dirs[dir])
			if _, ok := keypad[candidate]; ok {
				pos = candidate
			}
		}
		code.WriteRune(keypad[pos])
	}

	return code.String()
}

type direction rune

const (
	upDirection    direction = 'U'
	downDirection  direction = 'D'
	rightDirection direction = 'R'
	leftDirection  direction = 'L'
)
