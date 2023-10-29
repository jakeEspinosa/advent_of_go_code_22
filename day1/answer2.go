package main

import (
	"sort"
	"strconv"
)

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

func SortAndReturnTotal(intArray []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(intArray)))
	total := intArray[0] + intArray[1] + intArray[2]
	return total
}
