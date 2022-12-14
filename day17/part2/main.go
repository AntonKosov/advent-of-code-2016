package main

import (
	"crypto/md5"
	"fmt"

	"github.com/AntonKosov/advent-of-code-2016/aoc"
)

func main() {
	code := read()
	r := process(code)
	fmt.Printf("Answer: %v\n", r)
}

func read() string {
	return aoc.ReadAllInput()[0]
}

func process(code string) int {
	const size = 4
	target := aoc.NewVector2(size-1, size-1)
	current := []state{{pos: aoc.NewVector2(0, 0), code: code}}
	directions := []direction{
		{aoc.NewVector2(0, 1), "D"},
		{aoc.NewVector2(0, -1), "U"},
		{aoc.NewVector2(1, 0), "R"},
		{aoc.NewVector2(-1, 0), "L"},
	}

	longest := 0
	for len(current) > 0 {
		var next []state
		for _, s := range current {
			if s.pos == target {
				longest = len(s.code) - len(code)
				continue
			}
			openDoors := doorStates(s.code)
			for i, dir := range directions {
				if !openDoors[i] {
					continue
				}
				nextPos := s.pos.Add(dir.offset)
				if nextPos.X < 0 || nextPos.Y < 0 || nextPos.X >= size || nextPos.Y >= size {
					continue
				}
				next = append(next, state{pos: nextPos, code: s.code + dir.code})
			}
		}

		current = next
	}

	return longest
}

func doorStates(code string) [4]bool {
	hash := []byte(fmt.Sprintf("%x", md5.Sum([]byte(code))))
	isOpen := func(index int) bool { return hash[index] >= 'b' }
	return [4]bool{isOpen(1), isOpen(0), isOpen(3), isOpen(2)}
}

type state struct {
	pos  aoc.Vector2
	code string
}

type direction struct {
	offset aoc.Vector2
	code   string
}
