package main

import (
	"fmt"
)

func reversePrefix(word string, ch byte) string {
	wordByte := []byte(word)
	if len(word) == 1 {
		return word
	}

	var i, j int

	for num, l := range wordByte {
		if l == ch {
			j = num
			break
		}
	}

	if j == 0 {
		return word
	}

	for i < j {
		wordByte[i], wordByte[j] = wordByte[j], wordByte[i]
		i++
		j--
	}

	return string(wordByte)
}

func main() {
	fmt.Print(reversePrefix("abcdefd", 'd'))
}
