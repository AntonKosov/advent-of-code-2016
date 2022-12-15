package aoc

func Contains[T any](slice []T, equal func(value T) bool) bool {
	for _, v := range slice {
		if equal(v) {
			return true
		}
	}
	return false
}
