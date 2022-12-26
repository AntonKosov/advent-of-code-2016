package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Answer: %v\n", process())
}

func process() int {
	for a := 1; ; a++ {
		if try(a) {
			return a
		}
	}
}

func try(startA int) bool {
	digitsLeft := 100
	exValue := 0
	a, b, c, d := startA, 0, 0, 0
	// 00 cpy a d
	// 01 cpy 4 c
	// 02 cpy 643 b
	// 03 inc d
	// 04 dec b
	// 05 jnz b -2
	// 06 dec c
	// 07 jnz c -5
	d = a + 4*643
	for {
		// 08 cpy d a
		a = d
		for a != 0 {
			// 09 jnz 0 0
			// 10 cpy a b
			// 11 cpy 0 a
			b = a
			a = 0
		bLoop:
			for {
				// 12 cpy 2 c
				c = 2
				for {
					// 13 jnz b 2
					// 14 jnz 1 6
					if b == 0 {
						break bLoop
					}
					// 15 dec b
					// 16 dec c
					b--
					c--
					// 17 jnz c -4
					if c == 0 {
						break
					}
				}
				// 18 inc a
				a++
				// 19 jnz 1 -7
			}
			// 20 cpy 2 b
			// 21 jnz c 2
			// 22 jnz 1 4
			// 23 dec b
			// 24 dec c
			// 25 jnz 1 -4
			b = 2 - c // c = [1, 2]
			// c = 0 // don't need it
			// 26 jnz 0 0
			// 27 out b
			if exValue != b {
				return false
			}
			exValue = 1 - exValue
			digitsLeft--
			if digitsLeft == 0 {
				return true
			}
			// 28 jnz a -19
		}
		// 29 jnz 1 -21
	}
}
