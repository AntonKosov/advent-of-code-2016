package main

import (
	"fmt"
)

func main() {
	state := read()
	steps := process(state)
	fmt.Printf("Answer: %v\n", steps)
}

func read() state {
	return state{
		elevator: 0,
		devices: [4]byte{
			d0,
			d1 | d2,
			d3 | d4,
			0,
		},
		generators: [4]byte{
			d0 | d1 | d2,
			0,
			d3 | d4,
			0,
		},
	}
}

func process(startState state) int {
	targetState := state{
		elevator:   3,
		devices:    [4]byte{0, 0, 0, d0 | d1 | d2 | d3 | d4},
		generators: [4]byte{0, 0, 0, d0 | d1 | d2 | d3 | d4},
	}
	processedStates := map[state]bool{startState: true}
	steps := 0
	currentSteps := []state{startState}
	for !processedStates[targetState] {
		if len(currentSteps) == 0 {
			panic("something went wrong")
		}
		steps++
		var nextSteps []state
		for _, s := range currentSteps {
			generateNextSteps(s, processedStates, &nextSteps)
		}

		currentSteps = nextSteps
		fmt.Println(steps, len(currentSteps))
	}

	return steps
}

func generateNextSteps(currentState state, processedStates map[state]bool, nextSteps *[]state) {
	e := currentState.elevator
	devices := getIDs(currentState.devices[e])
	generators := getIDs(currentState.generators[e])
	moveToFloor := func(floor, devMask, genMask byte) {
		if floor >= countFloors {
			return
		}

		candidate := currentState
		candidate.elevator = floor
		candidate.devices[e] &^= devMask
		candidate.devices[floor] |= devMask
		candidate.generators[e] &^= genMask
		candidate.generators[floor] |= genMask

		if candidate.valid() && !processedStates[candidate] {
			processedStates[candidate] = true
			*nextSteps = append(*nextSteps, candidate)
		}
	}
	move := func(devMask, genMask byte) {
		moveToFloor(e-1, devMask, genMask)
		moveToFloor(e+1, devMask, genMask)
	}

	for i := 0; i < len(devices); i++ {
		d1 := devices[i]
		move(d1, 0)
		for j := i + 1; j < len(devices); j++ {
			d2 := devices[j]
			move(d1|d2, 0)
		}
		for _, gen := range generators {
			move(d1, gen)
		}
	}

	for i := 0; i < len(generators); i++ {
		g1 := generators[i]
		move(0, g1)
		for j := i + 1; j < len(generators); j++ {
			move(0, g1|generators[j])
		}
	}
}

func getIDs(floor byte) []byte {
	res := make([]byte, 0, countDevices)
	for i := 0; i < countDevices; i++ {
		d := byte(1 << i)
		if floor&d != 0 {
			res = append(res, d)
		}
	}

	return res
}

const (
	countDevices = 5
	countFloors  = 4
	// devices
	d0 byte = 1 << 0 // thulium
	d1 byte = 1 << 1 // plutonium
	d2 byte = 1 << 2 // strontium
	d3 byte = 1 << 3 // promethium
	d4 byte = 1 << 4 // ruthenium
)

// Floor plan:
//   - Devices
//     00054321
//   - Generators
//     00054321
type state struct {
	elevator   byte
	devices    [countFloors]byte
	generators [countFloors]byte
}

func (s state) valid() bool {
	for floor := 0; floor < countFloors; floor++ {
		genFloor := s.generators[floor]
		for device := 0; device < countDevices; device++ {
			mask := byte(1 << device)
			if s.devices[floor]&mask != 0 && genFloor&mask == 0 && genFloor&^mask != 0 {
				return false
			}
		}
	}

	return true
}
