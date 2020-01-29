package main

import (
	"fmt"
)

func main() {
	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	// fmt.Println("length: %d\n", length)
	// fmt.Println("input: %s\n", input)
	// fmt.Println("delta: %d\n", delta)

	alphabet := []rune("abcdefghijklmnopqrstuvwxyz")
	newRune := rotate('m', 2, alphabet)
	fmt.Println(string(newRune))
}

func rotate(s rune, delta int, key []rune) rune {
	idx := -1
	for i, r := range key {
		if r == s {
			idx = i
			break
		}
	}
	if idx < 0 {
		panic("idx < 0")
	}

	for i := 0; i < delta; i++ {
		idx++
		if idx >= len(key) {
			idx = 0
		}
	}

	return key[idx]
}
