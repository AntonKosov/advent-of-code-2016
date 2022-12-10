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

	triangles := make([]triangle, 0, len(lines))
	parse := func(column int) {
		from := column * 5
		to := from + 5
		for i := 0; i < len(lines); i += 3 {
			triangles = append(triangles, triangle{
				a: aoc.StrToInt(strings.TrimSpace(lines[i][from:to])),
				b: aoc.StrToInt(strings.TrimSpace(lines[i+1][from:to])),
				c: aoc.StrToInt(strings.TrimSpace(lines[i+2][from:to])),
			})
		}
	}

	parse(0)
	parse(1)
	parse(2)

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
