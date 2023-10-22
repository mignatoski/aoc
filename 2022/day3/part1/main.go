package part1

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

	var line string
	var sum int
	for fileScanner.Scan() {
		line = fileScanner.Text()
		sum += Prioritize(line)

		// fmt.Println(sum, line)
	}
	fmt.Println(sum)
}

func Prioritize(line string) int {
	var result int

	size := len(line) / 2
	c1 := []byte(line)[:size]
	c2 := []byte(line)[size:]

	rucksack := make(map[string]bool)

	for _, v := range c1 {
		rucksack[string(v)] = true
	}

	for _, v := range c2 {
		_, exists := rucksack[string(v)]
		if exists {
			result = PriorityList[string(v)]
			// fmt.Println(string(v), result)
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
