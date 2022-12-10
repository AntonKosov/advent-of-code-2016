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
			a.superSeq = append(a.superSeq, section(values[0]))
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
		var abas [][]byte
		for _, s := range addr.superSeq {
			abas = append(abas, s.aba()...)
		}
		for _, aba := range abas {
			a, b := aba[0], aba[1]
			for _, s := range addr.hyperSeq {
				if s.hasBAB(a, b) {
					count++
					continue nextAddr
				}
			}
		}
	}

	return count
}

type section []byte

func (s section) aba() [][]byte {
	var res [][]byte
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] && s[i] != s[i+1] {
			res = append(res, []byte(s[i:i+3]))
		}
	}

	return res
}

func (s section) hasBAB(a, b byte) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == b && s[i+1] == a && s[i+2] == b {
			return true
		}
	}

	return false
}

type address struct {
	superSeq []section
	hyperSeq []section
}
