package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2016/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() [][]string {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	commands := make([][]string, len(lines))
	for i, line := range lines {
		commands[i] = strings.Split(line, " ")
	}

	return commands
}

func process(commands [][]string) int {
	a := 12
	// 00 cpy a b
	// 01 dec b
	for b := a - 1; b > 1; b-- {
		// 02 cpy a d
		// 03 cpy 0 a
		// 04 cpy b c
		// 05 inc a
		// 06 dec c
		// 07 jnz c -2
		// 08 dec d
		// 09 jnz d -5
		// a = b * d
		a *= b
		// 10 dec b
		// 11 cpy b c
		// 12 cpy c d
		// 13 dec d
		// 14 inc c
		// 15 jnz d -2
		// 16 tgl c // c = b * 2
		// 17 cpy -16 c
		// 18 jnz 1 c // c == 2 => cpy
	}
	// 19 cpy 95 c
	// 20 jnz 91 d // c == 4 => cpy
	// 21 inc a
	// 22 inc d // c == 6 => dec
	// 23 jnz d -2
	// 24 inc c  // c == 8 => dec
	// 25 jnz c -5
	a += 91 * 95

	// a = 12! + 91 * 95
	return a
}
