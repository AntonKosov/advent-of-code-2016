package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2016/aoc"
)

func main() {
	r := process(
		aoc.NewVector2(1, 1),
		aoc.NewVector2(31, 39),
		1352,
	)
	fmt.Printf("Answer: %v\n", r)
}

func process(start, dest aoc.Vector2, favoriteNumber int) int {
	officeMap := newOffice(favoriteNumber)
	processed := map[aoc.Vector2]bool{}
	heap := aoc.NewHeap(func(v1, v2 cell) bool { return v1.score < v2.score })
	heap.Push(newCell(start, dest, 0))

	dirs := []aoc.Vector2{
		aoc.NewVector2(0, 1),
		aoc.NewVector2(0, -1),
		aoc.NewVector2(1, 0),
		aoc.NewVector2(-1, 0),
	}
	addCell := func(location aoc.Vector2, steps int) {
		if location.X < 0 || location.Y < 0 || !officeMap.empty(location) {
			return
		}
		heap.Push(newCell(location, dest, steps))
	}
	for heap.Len() > 0 {
		c := heap.Pop()
		if processed[c.location] {
			continue
		}
		processed[c.location] = true
		if c.location == dest {
			return c.steps
		}
		for _, dir := range dirs {
			addCell(c.location.Add(dir), c.steps+1)
		}
	}

	panic("path not found")
}

type cell struct {
	location aoc.Vector2
	steps    int
	score    int
}

func newCell(location, dest aoc.Vector2, steps int) cell {
	return cell{
		location: location,
		steps:    steps,
		score:    steps + dest.Sub(location).ManhattanDst(),
	}
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
