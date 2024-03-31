package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	Duration int
	Record   int
}

var (
	races []Race
)

func main() {
	// inputFile, _ := os.Open("sample.txt")
	inputFile, _ := os.Open("part2.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	fileScanner.Scan()
	line = fileScanner.Text()
	durations := strings.Fields(line)

	fileScanner.Scan()
	line = fileScanner.Text()
	records := strings.Fields(line)

	fmt.Println(durations, records)

	races = make([]Race, 0)
	for i := 1; i < len(durations); i++ {
		dur, _ := strconv.ParseInt(durations[i], 0, 0)
		rec, _ := strconv.ParseInt(records[i], 0, 0)
		r := Race{
			Duration: int(dur),
			Record:   int(rec),
		}
		fmt.Println(r)
		races = append(races, r)
	}

	part1 := 1
	for _, r := range races {
		part1 *= r.Ways()
	}

	fmt.Println("Part 1: ", part1)

	fmt.Println("Part 2: ", line)
}

func (r Race) Ways() int {
	ways := 0
	for i := 1; i < r.Duration; i++ {
		if r.Win(i) {
			ways++
		}
	}
	return ways
}

func (r Race) Win(c int) bool {
	// Dist = (T - C)*C = TC - C^2
	return (r.Duration*c)-(c*c) > r.Record
}
