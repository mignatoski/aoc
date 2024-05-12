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
	part1, part2 := 0, 0
	lineOffset := 0
	for fileScanner.Scan() {
		levels := make([][]int64, 0)
		line = fileScanner.Text()

		nums := strings.Split(line, " ")
		vals := make([]int64, len(nums))
		for i, n := range nums {
			vals[i], _ = strconv.ParseInt(n, 0, 0)
		}
		levels = append(levels, vals)
		level := 1
		for {
			fmt.Println(levels[level-1])
			end := true
			vals := make([]int64, len(levels[level-1])-1)
			levels = append(levels, vals)
			for i := 0; i < len(levels[level-1])-1; i++ {
				vals[i] = levels[level-1][i+1] - levels[level-1][i]
				if vals[i] != 0 {
					end = false
				}
			}
			if end {
				fmt.Println(levels[level])
				break
			}
			level++
		}
		var proj, prej int64
		for i := level; i > 0; i-- {
			proj += levels[i-1][len(levels[i-1])-1]
			prej = levels[i-1][0] - prej
		}
		fmt.Println(proj)
		part1 += int(proj)
		part2 += int(prej)

		lineOffset++
	}
	fmt.Println("Part 1: ", part1)

	fmt.Println("Part 2: ", part2)
}
