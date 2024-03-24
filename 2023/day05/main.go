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
	f, _ := os.Open("input.txt")
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

	rs := make(RangeSet, 0)
	for _, r := range seeds {
		rs = append(rs, r)
	}
	minval := expand(&rs, 0)
	fmt.Println(minval)

}

func expand(rs *RangeSet, layer int) int {
	fmt.Println("l:", layer)
	if layer == len(layers) {
		result := math.MaxInt
		for _, r := range *rs {
			result = min(result, r.Start)
		}
		return result
	}
	lrs := layers[layer]
	nrs := make(RangeSet, 0) // hold next level range set
	for len(*rs) > 0 {
		nr := (*rs)[0]
		*rs = (*rs)[1:] // remove current range
		for i := 0; i < len(lrs); i++ {
			lr := lrs[i]
			if nr.Overlap(lr) {
				var rrs RangeSet
				nr, rrs = nr.Split(lr)
				*rs = append(*rs, rrs...) // add remaining 0, 1, or 2 ranges to current layer
				break
			}
		}
		fmt.Println("nr:", nr)
		nrs = append(nrs, nr)
	}

	fmt.Println("nrs:", nrs)
	return expand(&nrs, layer+1)
}

func (r Range) End() int {
	return r.Start + r.Len - 1
}

func (r Range) Overlap(lr Range) bool {
	if (r.Start <= lr.Start && r.End() >= lr.Start) ||
		(lr.Start <= r.Start && lr.End() >= r.Start) {
		return true
	}
	return false
}

func (r Range) Split(lr Range) (nr Range, rrs RangeSet) {
	// assumes overlap = true
	fmt.Println("r:", r, "lr:", lr)
	start := max(r.Start, lr.Start)
	end := min(r.End(), lr.End())
	nr = Range{
		Start: start + lr.Offset,
		Len:   (end - start) + 1,
	}
	rrs = make(RangeSet, 0)
	if r.Start < start {
		rrs = append(rrs, Range{
			Start: r.Start,
			Len:   start - r.Start,
		})
	}
	if r.End() > end {
		rrs = append(rrs, Range{
			Start: end + 1,
			Len:   (r.End() - end),
		})
	}
	fmt.Println("rrs:", rrs)
	return nr, rrs
}
