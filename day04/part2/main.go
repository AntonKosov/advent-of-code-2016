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

func read() []room {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	rooms := make([]room, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, "[")
		parts = strings.Split(parts[0], "-")
		rooms[i] = room{
			name:     parts[:len(parts)-1],
			sectorID: aoc.StrToInt(parts[len(parts)-1]),
		}
	}

	return rooms
}

func process(rooms []room) {
	for _, r := range rooms {
		fmt.Println(r.decryptName(), r.sectorID)
	}
}

type room struct {
	name     []string
	sectorID int
}

func (r room) decryptName() string {
	var name strings.Builder
	for _, word := range r.name {
		if name.Len() > 0 {
			name.WriteRune(' ')
		}
		for _, l := range word {
			name.WriteRune('a' + rune((int(l-'a')+r.sectorID)%('z'-'a'+1)))
		}
	}

	return name.String()
}
