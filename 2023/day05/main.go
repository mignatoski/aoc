package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ValMap struct {
	DestStart, SourceStart, RangeLength int
}
type Almanac []*[]*ValMap

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	lines := 0
	seeds := make([]int, 0)
	almanac := make(Almanac, 0)
	var maps *[]*ValMap
	var curValMap *ValMap
	minVal := 9999999999999
	for fileScanner.Scan() {
		lines++
		line = fileScanner.Text()

		if lines == 1 {
			strSeeds := strings.Split(line, " ")
			for i := 1; i < len(strSeeds); i++ {
				intSeed, _ := strconv.ParseInt(strSeeds[i], 0, 0)
				seeds = append(seeds, int(intSeed))
			}
			continue
		}

		if line == "" {
			continue
		}

		if strings.Contains(line, ":") {
			m := make([]*ValMap, 0)
			maps = &m
			almanac = append(almanac, maps)
		} else {
			curValMap = &ValMap{}
			*maps = append(*maps, curValMap)
			fmt.Sscanf(line, "%d %d %d", &curValMap.DestStart, &curValMap.SourceStart, &curValMap.RangeLength)
		}
	}

	for _, s := range seeds {
		val := s
		fmt.Printf("s: %v\n", s)
		for _, arr := range almanac {
			for _, m := range *arr {
				if val >= m.SourceStart && val < m.SourceStart+m.RangeLength {
					val += m.DestStart - m.SourceStart
					break
				}
			}
		}
		fmt.Printf("val: %v\n", val)
		minVal = min(minVal, val)
	}

	fmt.Println("Part 1: ", minVal)

	fmt.Println("Part 2: ", line)
}
