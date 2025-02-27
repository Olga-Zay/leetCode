package main

import (
	"fmt"
)

func reverseString(s []byte) []byte {
	i := 0
	j := len(s) - 1

	for i < j {
		s[i], s[j] = s[j], s[i]

		i++
		j--
	}

	return s
}

func main() {
	fmt.Print(reverseString([]byte{72, 101, 108, 108, 111}))
}
