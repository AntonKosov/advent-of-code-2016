package main

import (
	"fmt"
)

func main() {
	// Solved Manually
	// The empty node should come to the goal cell (around nodes with big capacity).
	// After that, 5 moves are required to move the goal toward the target.
	fmt.Printf("Answer: %v\n", 2+7+12+13+5*33+1)
}
