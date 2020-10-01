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

	var ret []rune
	for _, ch := range input {
		ret = append(ret, cipher(ch, delta))
	}
	fmt.Println(string(ret))
}

func cipher(r rune, delta int) rune {
	if r >= 'A' && r <= 'Z' {
		return rotate(r, 'A', delta)
	}
	if r >= 'a' && r <= 'z' {
		return rotate(r, 'a', delta)
	}
	return r
}

func rotate(r rune, base, delta int) rune {
	tmp := int(r) - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}

// func rotate(s rune, delta int, key []rune) rune {
// 	idx := strings.IndexRune(string(key), s)

// 	// code below, is same as line above
// 	// for i, r := range key {
// 	// 	if r == s {
// 	// 		idx = i
// 	// 		break
// 	// 	}
// 	// }

// 	if idx < 0 {
// 		panic("idx < 0")
// 	}

// 	// using modulo % (remainder)
// 	idx = (idx + delta) % len(key)

// 	// using a forloop, could loop many times (as many items you have in alphabet)
// 	// for i := 0; i < delta; i++ {
// 	// 	idx++
// 	// 	if idx >= len(key) {
// 	// 		idx = 0
// 	// 	}
// 	// }

// 	return key[idx]
// }
