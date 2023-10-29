package main

import (
	"strconv"
)

// Strategy:
// sort the int arr in reverse order
// get the first 3 elements, sum them, and return

func ConvertStringArrayToIntArray(stringArray []string) []int {
	var intSlice []int
	for i := 0; i < len(stringArray); i++ {
		newInt, err := strconv.Atoi(stringArray[i])
		if err != nil {
			panic(err)
		}

		intSlice = append(intSlice, newInt)
	}

	return intSlice
}
