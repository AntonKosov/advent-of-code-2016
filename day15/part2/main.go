package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2016/aoc"
)

func main() {
	disks := read()
	r := process(disks)
	fmt.Printf("Answer: %v\n", r)
}

func read() []disk {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	disks := make([]disk, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		pos := parts[11]
		pos = pos[:len(pos)-1]
		disks[i] = disk{
			position:  aoc.StrToInt(pos),
			positions: aoc.StrToInt(parts[3]),
		}
	}

	return disks
}

func process(disks []disk) int {
	fd := disks[0]
nextTime:
	for time := fd.positions - fd.position - 1; ; time += fd.positions {
		for i := 1; i < len(disks); i++ {
			d := disks[i]
			pos := (d.position + time + i + 1) % d.positions
			if pos != 0 {
				continue nextTime
			}
		}
		return time
	}
}

type disk struct {
	position  int
	positions int
}
