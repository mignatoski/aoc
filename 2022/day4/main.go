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
	var count int
	for fileScanner.Scan() {
		line = fileScanner.Text()

		ranges := strings.Split(line, ",")
		sr1 := strings.Split(ranges[0], "-")
		sr2 := strings.Split(ranges[1], "-")

		r1 := make([]int64, 2)
		r2 := make([]int64, 2)

		r1[0], _ = strconv.ParseInt(sr1[0], 10, 0)
		r1[1], _ = strconv.ParseInt(sr1[1], 10, 0)

		r2[0], _ = strconv.ParseInt(sr2[0], 10, 0)
		r2[1], _ = strconv.ParseInt(sr2[1], 10, 0)

		// Full Overlap
		/* if (r1[0] <= r2[0] && r1[1] >= r2[1]) || (r1[0] >= r2[0] && r1[1] <= r2[1]) {
			count++
			fmt.Println(line, r1, r2, count)
		} */

		if (r1[0] <= r2[0] && r1[1] >= r2[0]) || (r1[0] <= r2[1] && r1[1] >= r2[1]) {
			count++
			fmt.Println(line, r1, r2, count)
		} else if (r2[0] <= r1[0] && r2[1] >= r1[0]) || (r2[0] <= r1[1] && r2[1] >= r1[1]) {
			count++
			fmt.Println(line, r1, r2, count)
		}

	}
	fmt.Println(count)
}
