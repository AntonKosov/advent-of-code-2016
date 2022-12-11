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

func read() map[int]*bot {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	bots := make(map[int]*bot, len(lines))
	botByID := func(id int) *bot {
		b := bots[id]
		if b == nil {
			b = &bot{}
			bots[id] = b
		}
		return b
	}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		switch op := parts[0]; op {
		case "value":
			value := aoc.StrToInt(parts[1])
			botID := aoc.StrToInt(parts[5])
			b := botByID(botID)
			b.chips = append(b.chips, value)
			if len(b.chips) > 2 {
				panic(fmt.Sprintf("too many chips: %+v", *b))
			}
		case "bot":
			b := botByID(aoc.StrToInt(parts[1]))
			b.lowerReceiver = receiver{
				receiverType: receiverType(parts[5]),
				id:           aoc.StrToInt(parts[6]),
			}
			b.higherReceiver = receiver{
				receiverType: receiverType(parts[10]),
				id:           aoc.StrToInt(parts[11]),
			}
		default:
			panic(fmt.Sprintf("unknown operation: %v", op))
		}
	}

	return bots
}

func process(bots map[int]*bot) int {
	var readyBots []int
	for id, b := range bots {
		if len(b.chips) == 2 {
			readyBots = append(readyBots, id)
		}
	}

	outputs := make([]*int, 3)
	outputsLeft := 3

	for len(readyBots) > 0 {
		id := readyBots[len(readyBots)-1]
		readyBots = readyBots[:len(readyBots)-1]
		b := bots[id]

		lower, higher := aoc.Min(b.chips[0], b.chips[1]), aoc.Max(b.chips[0], b.chips[1])
		pass := func(rec receiver, chip int) {
			switch rt := rec.receiverType; rt {
			case outputReceiver:
				outputID := rec.id
				if outputID < 3 && outputs[outputID] == nil {
					outputs[outputID] = &chip
					outputsLeft--
				}
			case botReceiver:
				recID := rec.id
				recBot := bots[recID]
				recBot.chips = append(recBot.chips, chip)
				if len(recBot.chips) == 2 {
					readyBots = append(readyBots, recID)
				}
			default:
				panic(fmt.Sprintf("unknown receiver: %v", rt))
			}
		}

		pass(b.lowerReceiver, lower)
		pass(b.higherReceiver, higher)
		b.chips = nil
		if outputsLeft == 0 {
			return *outputs[0] * *outputs[1] * *outputs[2]
		}
	}

	panic("no bots left")
}

type receiverType string

const (
	outputReceiver receiverType = "output"
	botReceiver    receiverType = "bot"
)

type receiver struct {
	receiverType receiverType
	id           int
}

type bot struct {
	chips          []int
	lowerReceiver  receiver
	higherReceiver receiver
}
