package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	for fileScanner.Scan() {
		line = fileScanner.Text()
	}
	fmt.Println(line)
}
