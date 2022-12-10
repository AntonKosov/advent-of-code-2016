package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2016/aoc"
)

func main() {
	data := read()
	process(data)
}

func read() []command {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	commands := make([]command, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		var cmd command
		switch c := parts[0]; c {
		case "rect":
			size := strings.Split(parts[1], "x")
			cmd = rectCommand{
				width:  aoc.StrToInt(size[0]),
				height: aoc.StrToInt(size[1]),
			}
		case "rotate":
			count := aoc.StrToInt(parts[4])
			source := aoc.StrToInt(strings.Split(parts[2], "=")[1])
			switch s := parts[1]; s {
			case "row":
				cmd = rotateRowCommand{
					row:   source,
					count: count,
				}
			case "column":
				cmd = rotateColumnCommand{
					column: source,
					count:  count,
				}
			default:
				panic(fmt.Sprintf("unknown source: %v", s))
			}
		default:
			panic(fmt.Sprintf("unknown command: %v", c))
		}
		commands[i] = cmd
	}

	return commands
}

func process(commands []command) {
	display := make([][]bool, 6)
	for i := range display {
		display[i] = make([]bool, 50)
	}

	for _, cmd := range commands {
		cmd.execute(display)
	}

	for _, row := range display {
		for _, lit := range row {
			var c rune
			if lit {
				c = '#'
			} else {
				c = '.'
			}
			fmt.Print(string(c))
		}
		fmt.Println()
	}
}

type command interface {
	execute(display [][]bool)
}

type rectCommand struct {
	width  int
	height int
}

func (c rectCommand) execute(display [][]bool) {
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			display[y][x] = true
		}
	}
}

type rotateColumnCommand struct {
	column int
	count  int
}

func (c rotateColumnCommand) execute(display [][]bool) {
	h := len(display)
	col := make([]bool, h)
	for i := 0; i < h; i++ {
		col[(i+c.count)%h] = display[i][c.column]
	}
	for i := 0; i < h; i++ {
		display[i][c.column] = col[i]
	}
}

type rotateRowCommand struct {
	row   int
	count int
}

func (c rotateRowCommand) execute(display [][]bool) {
	w := len(display[c.row])
	row := make([]bool, w)
	for i := 0; i < w; i++ {
		row[(i+c.count)%w] = display[c.row][i]
	}
	for i := 0; i < w; i++ {
		display[c.row][i] = row[i]
	}
}
