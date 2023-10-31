package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

/*
2D Grid of square nodes
Node has

Initialized:
- Neighbors
-- Up
-- Right
-- Down
-- Left
- Position (x,y)
- Height
- Type (Start, Normal, End)

Mutable:
- MinDistance

Calculated:
- DistanceToTarget
- CostToNeighbor
*/

type NodeType int

const (
	START NodeType = iota
	NORMAL
	END
)

type Node struct {
	Height, X, Y    int
	Letter          rune
	Type            NodeType
	GScore, FScore  int
	Neighbors       []*Node
	IsOpen, Visited bool
	Prev            *Node
}

func (n *Node) GetNeighbors() {
}

type Grid struct {
	Nodes          [][]Node
	Start, End     *Node
	sX, sY, eX, eY int
	Width, Height  int
}

func (g *Grid) SetNeighbors(n *Node) {
	// n := (*g).Nodes[y][x]
	x, y := n.X, n.Y
	if n.Neighbors == nil {
		n.Neighbors = make([]*Node, 0)
		if x-1 >= 0 {
			if (*g).Nodes[y][x-1].Height <= n.Height+1 {
				n.Neighbors = append(n.Neighbors, &(*g).Nodes[y][x-1])
			}
		}
		if x+1 < g.Width {
			if (*g).Nodes[y][x+1].Height <= n.Height+1 {
				n.Neighbors = append(n.Neighbors, &(*g).Nodes[y][x+1])
			}
		}
		if y-1 >= 0 {
			if (*g).Nodes[y-1][x].Height <= n.Height+1 {
				n.Neighbors = append(n.Neighbors, &(*g).Nodes[y-1][x])
			}
		}
		if y+1 < g.Height {
			if (*g).Nodes[y+1][x].Height <= n.Height+1 {
				n.Neighbors = append(n.Neighbors, &(*g).Nodes[y+1][x])
			}
		}
	}
}

func (g Grid) String() string {
	var s string
	for y := range g.Nodes {

		for x := range g.Nodes[y] {
			if g.Nodes[y][x].Letter > 'Z' && g.Nodes[y][x].Visited {
				s += "."
			} else {
				s += string(g.Nodes[y][x].Letter)
			}
		}
		s += "\n"
	}
	return s
}

func (g *Grid) H(n *Node) int {
	return int(math.Abs(float64(n.X-g.End.X)) + math.Abs(float64(n.Y-g.End.Y)))
}

var FinalPath []*Node

func (n *Node) PathBack() int {
	FinalPath = append(FinalPath, n)
	score := 0
	n.Letter = rune(strings.ToUpper(string(n.Letter))[0])
	if n.Prev != nil {
		score += 1 + n.Prev.PathBack()
	}
	fmt.Print("(", n.X, n.Y, ") -> ")
	return score

}

// func TravelCost(c *Node, n *Node) int {
// 	c.Height
// }

func (g *Grid) CalculatePath() {
	// start = g.Nodes[g.sY][g.sX]
	// end := g.Nodes[g.eY][g.eX]
	FinalPath = make([]*Node, 0)
	g.Start.FScore = g.H(g.Start)

	openSet := make([]*Node, 0)
	openSet = append(openSet, g.Start)

	for len(openSet) > 0 {
		sort.Slice(openSet, func(i, j int) bool {
			return openSet[i].FScore < openSet[j].FScore
		})
		// fmt.Println(openSet)

		current := openSet[0] // Lowest distance to start
		if current.X == g.End.X && current.Y == g.End.Y {
			fmt.Println("Score: ", current.PathBack())
			break
		}
		openSet = openSet[1:] // Remove current element
		current.IsOpen = false
		current.Visited = true

		g.SetNeighbors(current)

		for _, n := range current.Neighbors {
			gScoreTemp := current.GScore + 1
			if gScoreTemp < n.GScore {
				n.GScore = gScoreTemp
				n.FScore = gScoreTemp + g.H(n)
				n.Prev = current
				if !n.IsOpen {
					n.IsOpen = true
					openSet = append(openSet, n)
				}
			}
		}

	}

	// Start Node
	// set neighbors
	// add to unvisited list
	// sort by lowest optimistic distance

	// fmt.Println(g.Start, g.End)
	fmt.Println(len(FinalPath))

}

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	var y int
	grid := Grid{Nodes: make([][]Node, 0)}

	// Build Grid
	for fileScanner.Scan() {
		line = fileScanner.Text()

		grid.Nodes = append(grid.Nodes, make([]Node, 0))

		for x, r := range line {
			n := Node{Letter: r, GScore: math.MaxInt, FScore: math.MaxInt}

			switch r {
			case 'S':
				n.Type = START
				n.GScore = 0
				r = 'a'
				grid.sX, grid.sY = x, y
				grid.Start = &n
			case 'E':
				n.Type = END
				r = 'z'
				grid.eX, grid.eY = x, y
				grid.End = &n
			default:
				n.Type = NORMAL
			}
			n.Height = int(r - 'a')
			n.X, n.Y = x, y
			grid.Nodes[y] = append(grid.Nodes[y], n)
		}
		y++
	}
	grid.Width = len(grid.Nodes[0])
	grid.Height = len(grid.Nodes)

	// Start Processing Path
	grid.CalculatePath()
	fmt.Print(grid)

}
