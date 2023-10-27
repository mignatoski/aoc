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

	grid := make(Grid, 0)
	var line string
	var count int
	for fileScanner.Scan() {
		line = fileScanner.Text()
		grid = append(grid, Row{})
		for _, v := range []rune(line) {
			treeHeight, _ := strconv.ParseInt(string(v), 0, 0)
			grid[count] = append(grid[count], Tree(treeHeight))
		}
		count++
	}

	for _, r := range grid {
		fmt.Println(r)
	}

	numVisibile := 0
	for i, r := range grid {
		for j := range r {
			if IsVisible(&grid, j, i) {
				numVisibile++
				fmt.Print(1)
			} else {
				fmt.Print(0)
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("Visisble: ", numVisibile)

	maxScore := 0
	for i, r := range grid {
		for j := range r {
			score := Score(&grid, j, i)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	fmt.Println("Max Score: ", maxScore)

}

type Tree int
type Row []Tree
type Grid []Row
type Point struct{ x, y int }

func (g *Grid) Size() Point {
	return Point{len((*g)[0]), len(*g)}
}

// Checks tree at postion (x,y)
// (0,0) is top left corner, (0,n) is bottom left
func IsVisible(g *Grid, x, y int) bool {
	// Edges are always visible
	if x == 0 || y == 0 || x == g.Size().x-1 || y == g.Size().y-1 {
		return true
	}
	height := (*g)[y][x]
	n, s, e, w := true, true, true, true

	for i := 0; i < x; i++ {
		if (*g)[y][i] >= height {
			w = false
			break
		}
	}

	for i := x + 1; i < g.Size().x; i++ {
		if (*g)[y][i] >= height {
			e = false
			break
		}
	}

	for i := 0; i < y; i++ {
		if (*g)[i][x] >= height {
			n = false
			break
		}
	}

	for i := y + 1; i < g.Size().y; i++ {
		if (*g)[i][x] >= height {
			s = false
			break
		}
	}

	return n || s || e || w // if all checks pass assume visible
}

func Score(g *Grid, x, y int) int {
	// Edges are always visible
	if x == 0 || y == 0 || x == g.Size().x-1 || y == g.Size().y-1 {
		return 0
	}
	height := (*g)[y][x]
	n, s, e, w := 0, 0, 0, 0

	for i := x - 1; i >= 0; i-- {
		w++
		if (*g)[y][i] >= height {
			break
		}
	}

	for i := x + 1; i < g.Size().x; i++ {
		e++
		if (*g)[y][i] >= height {
			break
		}
	}

	for i := y - 1; i >= 0; i-- {
		n++
		if (*g)[i][x] >= height {
			break
		}
	}

	for i := y + 1; i < g.Size().y; i++ {
		s++
		if (*g)[i][x] >= height {
			break
		}
	}

	return n * s * e * w // if all checks pass assume visible
}
