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
	return aoc.ReadAllInput()[0]
}

func process(salt string) int {
	found := 0
	cache := map[int]hashInfo{}
	getHash := func(idx int) hashInfo {
		if h, ok := cache[idx]; ok {
			return h
		}
		str := fmt.Sprintf("%v%v", salt, idx)
		hi := newHashInfo(fmt.Sprintf("%x", md5.Sum([]byte(str))))
		cache[idx] = hi
		return hi
	}
nextHash:
	for i := 0; ; i++ {
		hash := getHash(i)
		if hash.firstThree != nil {
			three := *hash.firstThree
			for j := i + 1; j <= i+1000; j++ {
				h := getHash(j)
				if h.five[three] {
					found++
					fmt.Println(found, i)
					if found == 64 {
						return i
					}
					continue nextHash
				}
			}
		}
		delete(cache, i)
	}
}

type hashInfo struct {
	firstThree *byte
	five       map[byte]bool
}

func newHashInfo(hashStr string) hashInfo {
	hi := hashInfo{
		five: make(map[byte]bool, 0),
	}
	hash := []byte(hashStr)

	for i := 0; i < len(hash)-2; i++ {
		v := hash[i]
		if v == hash[i+1] && v == hash[i+2] {
			hi.firstThree = &v
			break
		}
	}

nextFive:
	for i := 0; i < len(hash)-4; i++ {
		v := hash[i]
		for j := i + 1; j < i+5; j++ {
			if v != hash[j] {
				i = j - 1
				continue nextFive
			}
		}
		hi.five[v] = true
	}

	return hi
}
