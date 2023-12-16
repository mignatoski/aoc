package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	var data, winning, mine []string
	dups := make(map[int]int)
	totalVal := 0
	numCards := 0
	for fileScanner.Scan() {
		line = fileScanner.Text()
		line = strings.ReplaceAll(line, "   ", " ")
		line = strings.ReplaceAll(line, "  ", " ")
		line = strings.ReplaceAll(line, ":", " :")
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
		id, _ := strconv.ParseInt(data[start-2], 0, 0)
		winning = data[start:sep]
		mine = data[sep+1:]

		val := 0
		winners := 0
		for i := range winning {
			for j := range mine {
				if winning[i] == mine[j] {
					winners++
					val = max(2*val, 1)
				}
			}
		}

		for i := int(id) + 1; i <= int(id)+winners; i++ {
			dups[i] += dups[int(id)] + 1
		}

		numCards += 1 + dups[int(id)]

		fmt.Println(winning, mine, val, winners)

		totalVal += val
	}
	fmt.Println("Part 1: ", totalVal)

	fmt.Println("Part 2: ", numCards)
}
