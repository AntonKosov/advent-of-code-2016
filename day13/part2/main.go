package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2016/aoc"
)

func main() {
	r := process(
		aoc.NewVector2(1, 1),
		1352,
	)
	fmt.Printf("Answer: %v\n", r)
}

func process(start aoc.Vector2, favoriteNumber int) int {
	officeMap := newOffice(favoriteNumber)
	visited := map[aoc.Vector2]bool{}
	dirs := []aoc.Vector2{
		aoc.NewVector2(0, 1),
		aoc.NewVector2(0, -1),
		aoc.NewVector2(1, 0),
		aoc.NewVector2(-1, 0),
	}
	current := []aoc.Vector2{start}
	for i := 0; i < 50; i++ {
		var next []aoc.Vector2
		for _, location := range current {
			for _, dir := range dirs {
				c := location.Add(dir)
				if c.X < 0 || c.Y < 0 || !officeMap.empty(c) || visited[c] {
					continue
				}
				visited[c] = true
				next = append(next, c)
			}
		}
		current = next
	}

	return len(visited)
}

type office struct {
	cache          map[aoc.Vector2]bool
	favoriteNumber int
}

func newOffice(favoriteNumber int) office {
	return office{
		cache:          map[aoc.Vector2]bool{},
		favoriteNumber: favoriteNumber,
	}
}

func (o *office) empty(location aoc.Vector2) bool {
	if v, ok := o.cache[location]; ok {
		return v
	}

	x, y := location.X, location.Y
	isEmpty := aoc.CountBits(x*x+3*x+2*x*y+y+y*y+o.favoriteNumber)&1 == 0
	o.cache[location] = isEmpty

	return isEmpty
}
