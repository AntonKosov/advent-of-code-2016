package aoc

func CountBits[T Ordinary](value T) int {
	count := 0
	for value != 0 {
		value &= value - 1
		count++
	}

	return count
}
