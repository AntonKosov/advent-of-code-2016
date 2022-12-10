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

func read() []triangle {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	triangles := make([]triangle, len(lines))
	for i, line := range lines {
		triangles[i] = triangle{
			a: aoc.StrToInt(strings.TrimSpace(line[:5])),
			b: aoc.StrToInt(strings.TrimSpace(line[5:10])),
			c: aoc.StrToInt(strings.TrimSpace(line[10:])),
		}
	}

	return triangles
}

func process(triangles []triangle) int {
	count := 0
	for _, t := range triangles {
		if t.possible() {
			count++
		}
	}

	return count
}

type triangle struct {
	a, b, c int
}

func (t triangle) possible() bool {
	return t.a+t.b > t.c && t.b+t.c > t.a && t.c+t.a > t.b
}
