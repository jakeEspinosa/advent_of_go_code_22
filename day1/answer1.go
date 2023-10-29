package main

import (
	"bufio"
	"fmt"
	"os"
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

func GetMostCalories(fileLines []string) int {
	largest := 0
	currentSum := 0
	for i := 0; i < len(fileLines); i++ {
		if fileLines[i] == "" {
			if largest < currentSum {
				largest = currentSum
			}
			currentSum = 0
		} else {
			lineAmount, err := strconv.Atoi(fileLines[i])
			if err != nil {
				panic(err)
			}
			currentSum = currentSum + lineAmount
		}
	}
	return largest
}

func main() {
	fmt.Printf("Reading input.txt...\n\n")
	lines := ReadFileLines("input.txt")
	mostCals := strconv.Itoa(GetMostCalories(lines))
	fmt.Printf("The elf carrying the most calories has:\n" + mostCals + " calories\n\n")
}
