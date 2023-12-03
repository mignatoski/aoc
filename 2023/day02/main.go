package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputFile, _ := os.Open("sample.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	for fileScanner.Scan() {
		line = fileScanner.Text()
	}
	fmt.Println("Part 1: ", line)

	fmt.Println("Part 2: ", line)
}
