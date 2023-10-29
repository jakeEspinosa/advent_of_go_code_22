package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func ReadFileLines(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	file.Close()
	return fileLines
}

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

func GetMostCalories(intArray []int) int {
	return intArray[1]
}

func main() {
	fmt.Printf("Reading input.txt...\n\n")

	lines := ReadFileLines("input.txt")
	sums := SumStringArray(lines)

	mostCals := strconv.Itoa(GetMostCalories(sums))
	fmt.Printf("The elf carrying the most calories has:\n" + mostCals + " calories\n\n")

	topThreeTotal := strconv.Itoa(SortAndReturnTotal(sums))
	fmt.Printf("The total of the top three elves is:\n" + topThreeTotal + "\n")
}
