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

func read() []instruction {
	lines := aoc.ReadAllInput()
	lines = strings.Split(lines[0], ", ")

	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		instructions[i] = instruction{
			dir:    direction(line[0]),
			blocks: aoc.StrToInt(line[1:]),
		}
	}

	return instructions
}

func process(instructions []instruction) int {
	pos := aoc.NewVector2(0, 0)
	dir := aoc.NewVector2(0, 1)
	visited := map[aoc.Vector2]bool{pos: true}
	for _, inst := range instructions {
		switch d := inst.dir; d {
		case leftDirection:
			dir = dir.RotateLeft()
		case rightDirection:
			dir = dir.RotateRight()
		default:
			panic(fmt.Sprintf("unknown direction: %v", d))
		}
		for i := 0; i < inst.blocks; i++ {
			pos = pos.Add(dir)
			if visited[pos] {
				return pos.ManhattanDst()
			}
			visited[pos] = true
		}
	}

	panic("invalid input")
}

type instruction struct {
	dir    direction
	blocks int
}

type direction rune

const (
	leftDirection  direction = 'L'
	rightDirection direction = 'R'
)
