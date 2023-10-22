package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	fmt.Printf("Reading input.txt...\n")
	lines := ReadFileLines("input.txt")
	for i := 0; i < len(lines); i++ {
		fmt.Printf(lines[i] + "\n")
	}
}
