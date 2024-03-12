package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SeedRange struct {
	Start, Length int
}
type ValMap struct {
	DestStart, SourceStart, RangeLength int
}
type ConversionTable []ValMap
type Almanac []ConversionTable

func main() {
	inputFile, _ := os.Open("sample.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	lines := 0
	seeds := make([]int, 0)
	seedRanges := make([]SeedRange, 0)
	almanac := make(Almanac, 0)
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
			for i := 1; i < len(strSeeds); i++ {
				intStart, _ := strconv.ParseInt(strSeeds[i], 0, 0)
				i++
				intLength, _ := strconv.ParseInt(strSeeds[i], 0, 0)
				seedRanges = append(seedRanges, SeedRange{int(intStart), int(intLength)})
			}
			continue
		}

		if line == "" {
			continue
		}

		if strings.Contains(line, ":") {
			m := make(ConversionTable, 0)
			almanac = append(almanac, m)
		} else {
			vm := ValMap{}
			fmt.Sscanf(line, "%d %d %d", &vm.DestStart, &vm.SourceStart, &vm.RangeLength)
			almanac[len(almanac)-1] = append(almanac[len(almanac)-1], vm)
		}
	}

	for _, s := range seeds {
		val := s
		fmt.Printf("s: %v\n", s)
		for _, arr := range almanac {
			for _, m := range arr {
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

	fmt.Println(seedRanges)
	fmt.Println(almanac)

	for i, x := range almanac {
		fmt.Println("Level:", i)
		for _, y := range x {
			fmt.Println(y)
		}
	}
	// for _, s := range seedRanges {
	// 	val := s
	// 	fmt.Printf("s: %v\n", s)
	// 	for _, arr := range almanac {
	// 		for _, m := range *arr {
	// 			if val >= m.SourceStart && val < m.SourceStart+m.RangeLength {
	// 				val += m.DestStart - m.SourceStart
	// 				break
	// 			}
	// 		}
	// 	}
	// 	fmt.Printf("val: %v\n", val)
	// 	minVal = min(minVal, val)
	// }
	//
	fmt.Println("Part 2: ", line)
}
