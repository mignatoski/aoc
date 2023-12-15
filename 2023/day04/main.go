package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	var data, winning, mine []string
	totalVal := 0
	for fileScanner.Scan() {
		line = fileScanner.Text()
		line = strings.ReplaceAll(line, "   ", " ")
		line = strings.ReplaceAll(line, "  ", " ")
		data = strings.Split(line, " ")

		start, sep := 0, 0
		for i, s := range data {
			switch s {
			case ":":
				start = i + 1
			case "|":
				sep = i
			}
		}
		winning = data[start:sep]
		mine = data[sep+1:]

		val := 0
		for i := range winning {
			for j := range mine {
				if winning[i] == mine[j] {
					val = max(2*val, 1)
				}
			}
		}
		fmt.Println(winning, mine, val)
		totalVal += val
	}
	fmt.Println("Part 1: ", totalVal)

	fmt.Println("Part 2: ", line)
}
