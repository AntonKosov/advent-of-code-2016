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
	cpu := [4]int{0, 0, 1, 0}
	i := 0
	parseRegister := func(str string) byte { return str[0] - 'a' }
	getValue := func(str string) int {
		if str >= "a" && str <= "d" {
			return cpu[parseRegister(str)]
		}
		return aoc.StrToInt(str)
	}
	handlers := map[string]func(command []string){
		"cpy": func(command []string) { cpu[parseRegister(command[2])] = getValue(command[1]) },
		"inc": func(command []string) { cpu[parseRegister(command[1])]++ },
		"dec": func(command []string) { cpu[parseRegister(command[1])]-- },
		"jnz": func(command []string) {
			if getValue(command[1]) != 0 {
				i += aoc.StrToInt(command[2]) - 1
			}
		},
	}

	for i < len(commands) {
		command := commands[i]
		i++
		handlers[command[0]](command)
	}

	return cpu[0]
}
