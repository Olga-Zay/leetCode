package main

import (
	"fmt"
	"strings"
)

func literalOrNumber(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func isPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}

	i := 0
	j := len(s) - 1

	for i < j {
		if !literalOrNumber(s[i]) {
			i++
			continue
		}

		if !literalOrNumber(s[j]) {
			j--
			continue
		}

		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		}

		i++
		j--
	}

	return true
}

func main() {
	fmt.Print(isPalindrome("A man, a plan, a canal: Panama"))
}
