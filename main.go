package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1

	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}

		if sum > target {
			j--
		}

		if sum < target {
			i++
		}
	}

	return nil
}

func main() {
	fmt.Print(twoSum([]int{2, 7, 11, 15}, 9))
}
