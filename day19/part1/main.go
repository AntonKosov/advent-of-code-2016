package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2016/aoc"
)

func main() {
	gamers := read()
	r := process(gamers)
	fmt.Printf("Answer: %v\n", r)
}

func read() int {
	return aoc.StrToInt(aoc.ReadAllInput()[0])
}

func process(gamers int) int {
	currentPlayer := buildList(gamers)
	for currentPlayer.next != currentPlayer {
		next := currentPlayer.next
		currentPlayer.presents += next.presents
		currentPlayer.next = next.next
		currentPlayer = currentPlayer.next
	}

	return currentPlayer.number
}

func buildList(gamers int) *node {
	first := &node{number: 1, presents: 1}
	tail := first
	for i := 1; i < gamers; i++ {
		tail.next = &node{number: i + 1, presents: 1}
		tail = tail.next
	}
	tail.next = first

	return first
}

type node struct {
	number   int
	presents int
	next     *node
}
