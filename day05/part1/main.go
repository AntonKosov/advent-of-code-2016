package main

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2016/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() string {
	lines := aoc.ReadAllInput()

	return lines[0]
}

func process(roomID string) string {
	var password strings.Builder
	for i := 0; password.Len() < 8; i++ {
		str := []byte(fmt.Sprintf("%v%v", roomID, i))
		hash := md5.Sum(str)
		if hash[0] == 0 && hash[1] == 0 && hash[2]&0xf0 == 0 {
			ch := fmt.Sprintf("%x", hash)
			password.WriteByte(ch[5])
		}
	}

	return password.String()
}
