package main

import (
	"crypto/md5"
	"fmt"

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
	const length = 8
	const notFound byte = 0xff

	password := make([]byte, length)
	for i := range password {
		password[i] = notFound
	}

	count := 0
	for i := 0; count < length; i++ {
		str := []byte(fmt.Sprintf("%v%v", roomID, i))
		hash := md5.Sum(str)
		if hash[0] == 0 && hash[1] == 0 && hash[2]&0xf0 == 0 {
			hashStr := fmt.Sprintf("%x", hash)
			index := int(hashStr[5] - '0')
			if index >= 0 && index < length && password[index] == notFound {
				password[index] = hashStr[6]
				count++
				fmt.Println(string(password))
			}
		}
	}

	return string(password)
}
