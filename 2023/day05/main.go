package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Range struct {
	Start, Offset, Len int
}
type RangeSet []Range
type Layers []RangeSet

var (
	seeds  RangeSet
	layers Layers
)

func main() {
	f, _ := os.Open("sample.txt")
	s := bufio.NewScanner(f)

	seeds = make(RangeSet, 0)
	layers = make(Layers, 0)
	li := -1

	for s.Scan() {
		l := s.Text()

		switch {
		case l == "":
		case strings.Contains(l, "seeds:"):
			nums := strings.Split(l, " ")
			nums = nums[1:]
			for i := 0; i < len(nums); i = i + 2 {
				var start, length int
				fmt.Sscanf(nums[i], "%d", &start)
				fmt.Sscanf(nums[i+1], "%d", &length)
				seeds = append(seeds, Range{start, 0, length})
			}
		case strings.Contains(l, "map:"):
			li++
			rs := make(RangeSet, 0)
			layers = append(layers, rs)
		default:
			nums := strings.Split(l, " ")
			var dest, src, length int
			fmt.Sscanf(nums[0], "%d", &dest)
			fmt.Sscanf(nums[1], "%d", &src)
			fmt.Sscanf(nums[2], "%d", &length)
			layers[li] = append(layers[li], Range{src, dest - src, length})
		}
	}

	fmt.Println(seeds[0], seeds[1])
	fmt.Println(layers[0][0], layers[1][0])

	for _, r := range seeds {
		sl := make(Layers, 0)
		rs := make(RangeSet, 0)
		rs = append(rs, r)
		sl = append(sl, rs)
		minval := expand(&sl, 0)
		fmt.Println(minval)
	}

}

func expand(sl *Layers, layer int) int {
	rs := layers[layer]
	if layer == len(layers)-1 {
		result := math.MaxInt
		for _, r := range rs {
			result = min(result, r.Start)
		}
		return result
	}
	for i := 0; i < len(rs); i++ {
		// remove current range
		for _, l := range layers {
			// if overlap
			// then add overlap to next layer
			//      add remaining 0, 1, or 2 ranges to current layer
			//      break
			// else add current range to next layer
			fmt.Println(l)
		}
	}

	return expand(sl, layer+1)
}
