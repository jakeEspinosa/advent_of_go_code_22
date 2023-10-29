package main

import (
	"strconv"
)

// Strategy:
// sort the int arr in reverse order
// get the first 3 elements, sum them, and return

func SumStringArray(stringArray []string) []int {
	var result []int

	currentSum := 0
	for i := 0; i < len(stringArray); i++ {
		if stringArray[i] == "" {
			result = append(result, currentSum)
			currentSum = 0

		} else {
			lineAmount, err := strconv.Atoi(stringArray[i])
			if err != nil {
				panic(err)
			}
			currentSum += lineAmount
		}
	}

	return result
}
