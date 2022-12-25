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
	cpu := [4]int{7, 0, 0, 0}
	i := 0
	parseRegister := func(str string) (byte, bool) {
		if str < "a" || str > "d" {
			return 0, false
		}
		return str[0] - 'a', true
	}
	getValue := func(str string) int {
		if str >= "a" && str <= "d" {
			r, _ := parseRegister(str)
			return cpu[r]
		}
		return aoc.StrToInt(str)
	}
	writeToRegister := func(register string, value int) {
		r, ok := parseRegister(register)
		if !ok {
			return
		}
		cpu[r] = value
	}

	for i >= 0 && i < len(commands) {
		command := commands[i]
		i++
		switch cmd := command[0]; cmd {
		case cpyCommand:
			writeToRegister(command[2], getValue(command[1]))
		case incCommand:
			writeToRegister(command[1], getValue(command[1])+1)
		case decCommand:
			writeToRegister(command[1], getValue(command[1])-1)
		case jnzCommand:
			if getValue(command[1]) != 0 {
				i += getValue(command[2]) - 1
			}
		case tglCommand:
			idx := i + getValue(command[1]) - 1
			if idx < 0 || idx >= len(commands) {
				continue
			}
			chCommand := commands[idx]
			if len(chCommand) == 2 {
				if chCommand[0] == incCommand {
					chCommand[0] = decCommand
				} else {
					chCommand[0] = incCommand
				}
				continue
			}
			if chCommand[0] == jnzCommand {
				chCommand[0] = cpyCommand
			} else {
				chCommand[0] = jnzCommand
			}
		}
	}

	return getValue("a")
}

const (
	cpyCommand = "cpy"
	incCommand = "inc"
	decCommand = "dec"
	jnzCommand = "jnz"
	tglCommand = "tgl"
)
