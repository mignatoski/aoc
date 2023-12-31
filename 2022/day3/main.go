package main

import (
	"bufio"
	"fmt"
	"os"
)

var PriorityList map[string]int

func init() {
	PriorityList = NewPriorityList()
}

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	line := make([]string, 3)
	var sum int
	for fileScanner.Scan() {
		line[0] = fileScanner.Text()
		fileScanner.Scan()
		line[1] = fileScanner.Text()
		fileScanner.Scan()
		line[2] = fileScanner.Text()
		sum += Prioritize(line)

		fmt.Println(sum, line)
	}
	fmt.Println(sum)
}

func Prioritize(line []string) int {
	var result int

	rucksack := make(map[string]bool)

	for _, v := range line[0] {
		rucksack[string(v)] = false
	}

	for _, v := range line[1] {
		_, exists := rucksack[string(v)]
		if exists {
			rucksack[string(v)] = true
		}
	}

	for _, v := range line[2] {
		inBoth, exists := rucksack[string(v)]
		if exists && inBoth {
			result = PriorityList[string(v)]
			fmt.Println(string(v), result)
		}
	}

	return result
}

func NewPriorityList() map[string]int {
	result := make(map[string]int, 52)

	for i, l := 1, 'a'; i <= 26; i++ {
		result[string(l)] = i
		l++
	}

	for i, l := 27, 'A'; i <= 52; i++ {
		result[string(l)] = i
		l++
	}

	return result
}
