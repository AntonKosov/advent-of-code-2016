package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2016/aoc"
)

const input = "abcdefgh"

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []operation {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	operations := make([]operation, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		var op operation
		switch verb := parts[0]; verb {
		case "rotate":
			switch dir := parts[1]; dir {
			case "left":
				count := aoc.StrToInt(parts[2])
				op = func(input []byte) []byte { return rotateLeft(input, count) }
			case "right":
				count := aoc.StrToInt(parts[2])
				op = func(input []byte) []byte { return rotateRight(input, count) }
			case "based":
				letter := parts[6][0]
				op = func(input []byte) []byte { return rotateBasedOnLetterPosition(input, letter) }
			default:
				panic(fmt.Sprintf("unknown direction: %v", dir))
			}
		case "swap":
			switch item := parts[1]; item {
			case "letter":
				from, to := parts[2][0], parts[5][0]
				op = func(input []byte) []byte { return swapLetter(input, from, to) }
			case "position":
				from, to := aoc.StrToInt(parts[2]), aoc.StrToInt(parts[5])
				op = func(input []byte) []byte { return swapPosition(input, from, to) }
			default:
				panic(fmt.Sprintf("unknown item: %v", item))
			}
		case "reverse":
			from, to := aoc.StrToInt(parts[2]), aoc.StrToInt(parts[4])
			op = func(input []byte) []byte { return reverse(input, from, to) }
		case "move":
			from, to := aoc.StrToInt(parts[2]), aoc.StrToInt(parts[5])
			op = func(input []byte) []byte { return movePosition(input, from, to) }
		default:
			panic(fmt.Sprintf("unknown verb: %v", verb))
		}
		operations[i] = op
	}

	return operations
}

func process(operations []operation) string {
	psw := []byte(input)
	for _, op := range operations {
		psw = op(psw)
	}

	return string(psw)
}

type operation func(input []byte) []byte

func rotateLeft(input []byte, count int) []byte {
	n := len(input)
	output := make([]byte, n)
	for i := range output {
		idx := (i + count) % n
		output[i] = input[idx]
	}
	return output
}

func rotateRight(input []byte, count int) []byte {
	n := len(input)
	output := make([]byte, n)
	for i := range output {
		idx := (i - count) % n
		if idx < 0 {
			idx += n
		}
		output[i] = input[idx]
	}
	return output
}

func rotateBasedOnLetterPosition(input []byte, letter byte) []byte {
	idx := 0
	for ; input[idx] != letter; idx++ {
	}
	count := idx + 1
	if idx >= 4 {
		count++
	}
	return rotateRight(input, count)
}

func swapLetter(input []byte, from, to byte) []byte {
	output := make([]byte, len(input))
	for i, r := range input {
		var tv byte
		switch r {
		case from:
			tv = to
		case to:
			tv = from
		default:
			tv = r
		}
		output[i] = tv
	}
	return output
}

func swapPosition(input []byte, from, to int) []byte {
	output := make([]byte, len(input))
	copy(output, input)
	output[from], output[to] = output[to], output[from]
	return output
}

func reverse(input []byte, from, to int) []byte {
	output := make([]byte, len(input))
	copy(output, input)
	for i := from; i <= to; i++ {
		output[i] = input[to+from-i]
	}
	return output
}

func movePosition(input []byte, from, to int) []byte {
	tmp := make([]byte, 0, len(input)-1)
	tmp = append(tmp, input[:from]...)
	tmp = append(tmp, input[from+1:]...)
	output := make([]byte, 0, len(input))
	output = append(output, tmp[:to]...)
	output = append(output, input[from])
	output = append(output, tmp[to:]...)
	return output
}
