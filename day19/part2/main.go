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
	looser := currentPlayer
	for i := 0; i < gamers/2; i++ {
		looser = looser.next
	}

	for currentPlayer.next != currentPlayer {
		currentPlayer.presents += looser.presents
		nextLooser := looser.next
		if gamers&1 == 1 {
			nextLooser = nextLooser.next
		}
		looser.prev.next = looser.next
		looser.next.prev = looser.prev
		looser = nextLooser
		currentPlayer = currentPlayer.next
		gamers--
	}

	return currentPlayer.number
}

func buildList(gamers int) *node {
	first := &node{number: 1, presents: 1}
	tail := first
	for i := 1; i < gamers; i++ {
		tail.next = &node{number: i + 1, presents: 1, prev: tail}
		tail = tail.next
	}
	tail.next = first
	first.prev = tail

	return first
}

type node struct {
	number   int
	presents int
	prev     *node
	next     *node
}
