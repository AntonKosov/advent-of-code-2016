package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/AntonKosov/advent-of-code-2016/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() []room {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	rooms := make([]room, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, "[")
		checksum := parts[1][:len(parts[1])-1]
		parts = strings.Split(parts[0], "-")
		name := strings.Join(parts[:len(parts)-1], "")
		rooms[i] = room{
			name:     name,
			checksum: checksum,
			sectorID: aoc.StrToInt(parts[len(parts)-1]),
		}
	}

	return rooms
}

func process(rooms []room) int {
	count := 0
	for _, r := range rooms {
		if r.valid() {
			count += r.sectorID
		}
	}

	return count
}

type room struct {
	name     string
	checksum string
	sectorID int
}

func (r room) valid() bool {
	freq := map[rune]int{}
	for _, r := range r.name {
		freq[r]++
	}
	list := make([]rune, 0, len(freq))
	for r := range freq {
		list = append(list, r)
	}
	sort.Slice(list, func(i, j int) bool {
		ri, rj := list[i], list[j]
		fi, fj := freq[ri], freq[rj]
		if fi == fj {
			return ri < rj
		}
		return fi > fj
	})

	cs := []rune(r.checksum)
	for i := 0; i < 5; i++ {
		if list[i] != cs[i] {
			return false
		}
	}

	return true
}
