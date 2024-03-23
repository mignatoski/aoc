package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SeedRange struct {
	Start, Length int
}
type ValMap struct {
	DestStart, SourceStart, RangeLength, MinOutput int
}
type ConversionTable []ValMap
type Almanac []ConversionTable

type ByMinOutput ConversionTable

func (a ByMinOutput) Len() int           { return len(a) }
func (a ByMinOutput) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByMinOutput) Less(i, j int) bool { return a[i].MinOutput < a[j].MinOutput }

type ByDestStart ConversionTable

func (a ByDestStart) Len() int           { return len(a) }
func (a ByDestStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDestStart) Less(i, j int) bool { return a[i].DestStart < a[j].DestStart }

type Graph struct {
	Nodes []*Node
}

type Node struct {
	IsRoot, IsLeaf     bool
	R                  Range
	Offset             int
	ToNodes, FromNodes []*Node
}

type Range struct {
	Start, End int
}

// Util

func Overlap(r1, r2 Range) bool {
	switch {
	case r1.Start >= r2.Start && r1.Start <= r2.End:
		return true
	case r1.End >= r2.Start && r1.End <= r2.End:
		return true
	default:
		return false
	}
}

func (n *Node) LinkTo(to *Node) {
	n.ToNodes = append(n.ToNodes, to)
	to.FromNodes = append(to.FromNodes, n)
}

func (g *Graph) DFSFindMin(r Range) (result int) {
	result = math.MaxInt
	for _, n := range g.Nodes {
		if n.IsRoot {
			result = min(result, n.FindMin(r))
		}
	}
	return result
}

func (n *Node) FindMin(r Range) (result int) {
	for _, t := range n.ToNodes {
		if Overlap(r, t.R) {
			if t.IsLeaf {
				result = max(r.Start, t.R.Start) + t.Offset
			} else {
				result = t.FindMin(Range{max(r.Start, t.R.Start), min(r.End, t.R.End)})
			}
		}
	}
	return result
}

// End Util

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
			i := len(almanac) - 1
			almanac[i] = append(almanac[i], vm)
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

	for i := len(almanac); i > 0; i-- {
		a := &almanac[i-1]
		if i == len(almanac) {
			sort.Sort(ByDestStart(*a))
		} else {
			sort.Sort(ByMinOutput(*a))
		}
		for j := range *a {
			if i == len(almanac) {
				(*a)[j].MinOutput = (*a)[j].DestStart
			} else {
				pmin := math.MaxInt
				um := math.MaxInt
				for _, c := range almanac[i] {
					if (*a)[j].DestStart+(*a)[j].RangeLength > c.SourceStart &&
						c.SourceStart+c.RangeLength > (*a)[j].DestStart {
						pmin = min(pmin, c.MinOutput)
					} else {
						if um < c.SourceStart {
							um = c.SourceStart - 1
						}

					}
				}
				(*a)[j].MinOutput = min(pmin, um)
			}
		}
	}

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
