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

func read() []address {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	addresses := make([]address, len(lines))
	for i, line := range lines {
		a := address{}
		parts := strings.Split(line, "]")
		for _, pair := range parts {
			values := strings.Split(pair, "[")
			a.seq = append(a.seq, section(values[0]))
			if len(values) > 1 {
				a.hyperSeq = append(a.hyperSeq, section(values[1]))
			}
		}
		addresses[i] = a
	}

	return addresses
}

func process(addresses []address) int {
	count := 0
nextAddr:
	for _, addr := range addresses {
		for _, s := range addr.hyperSeq {
			if s.hasABBA() {
				continue nextAddr
			}
		}
		for _, s := range addr.seq {
			if s.hasABBA() {
				count++
				break
			}
		}
	}

	return count
}

type section []byte

func (s section) hasABBA() bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] == s[i+3] && s[i+1] == s[i+2] && s[i] != s[i+1] {
			return true
		}
	}

	return false
}

type address struct {
	seq      []section
	hyperSeq []section
}
