package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var currentTotal int64
	var maxArr [3]int64

	var line string
	for fileScanner.Scan() {
		line = fileScanner.Text()
		if line != "" {
			lineInt, _ := strconv.ParseInt(line, 10, 0)
			currentTotal += lineInt
		} else {
			updateMaxArray(&maxArr, currentTotal)
			currentTotal = 0
		}
	}
	updateMaxArray(&maxArr, currentTotal)
	fmt.Println(maxArr)

	var sum int64
	for _, v := range maxArr {
		sum += v
	}
	fmt.Println(sum)
}

func updateMaxArray(maxArr *[3]int64, i int64) {
	switch true {
	case maxArr[0] < i:
		maxArr[2] = maxArr[1]
		maxArr[1] = maxArr[0]
		maxArr[0] = i

	case maxArr[1] < i:
		maxArr[2] = maxArr[1]
		maxArr[1] = i

	case maxArr[2] < i:
		maxArr[2] = i
	}
}
