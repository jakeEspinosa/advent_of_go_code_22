package main

// import (
// 	"strconv"
// )

// Strategy:
// convert the string arr to an int arr
// sort the int arr in reverse order
// get the first 3 elements, sum them, and return

// func GetTopThree(fileLines []string) int {
// 	topThreeArray := [3]int{0, 0, 0}

// 	currentSum := 0
// 	for i := 0; i < len(fileLines); i++ {
// 		if fileLines[i] == "" {
// 			if topThreeArray[0] < currentSum {
// 				topThreeArray[0] = currentSum
// 			} else if topThreeArray[1] < currentSum {
// 				topThreeArray[1] = currentSum
// 			} else if topThreeArray[2] < currentSum {
// 				topThreeArray[2] = currentSum
// 			}
// 			currentSum = 0

// 		} else {
// 			lineAmount, err := strconv.Atoi(fileLines[i])
// 			if err != nil {
// 				panic(err)
// 			}
// 			currentSum = currentSum + lineAmount
// 		}
// 	}
// 	return topThreeArray[0] + topThreeArray[1] + topThreeArray[2]
// }
