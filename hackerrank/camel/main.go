package main

import "fmt"

import "strings"

func main() {
	var input string
	fmt.Scanf("%s\n", &input)

	answer := 1
	for _, ch := range input {

		// OPTION 1:
		str := string(ch)
		if strings.ToUpper(str) == str {
			answer++
		}

		// OPTION 2:
		// min, max := 'A', 'Z'
		// if ch >= min && ch <= max {
		// 	// is a capital letter
		// 	answer++
		// }
	}
	fmt.Println(answer)
}
