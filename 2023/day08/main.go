package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Name, Value string
	L, R        *Node
}

var (
	nodes map[string]*Node
	cnt   int
	lr    []rune
)

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string

	fileScanner.Scan()
	lr = []rune(fileScanner.Text())

	fileScanner.Scan() // read blank line

	nodes = make(map[string]*Node, 0)
	curs := make([]*Node, 0)
	for fileScanner.Scan() {
		line = fileScanner.Text()
		n := Node{}
		n.Name, n.Value, _ = strings.Cut(line, " = ")
		nodes[n.Name] = &n
	}

	for _, n := range nodes {
		var l, r string
		fmt.Sscanf(
			strings.Replace(
				strings.Replace(
					strings.Replace(n.Value, ",", "", 1),
					"(", "", 1), ")", "", 1),
			"%v %v", &l, &r)
		n.L = nodes[l]
		n.R = nodes[r]
		// fmt.Println(l, r, n.L, n.R)

		if n.Name[2] == 'A' {
			curs = append(curs, n)
		}
	}

	cnt := len(lr)
	// fmt.Println(lr, nodes)
	// part1()

	fmt.Println(len(curs))

	cycles := make(map[State]int)
	nodeStepArray := make(map[*Node][]int)
	var part2 uint64 = 1
	for j, cur := range curs {
		steps := 0
		stepArray := make([]int, 0)
		nodeStepArray[cur] = stepArray
		for i, end := 0, false; !end; i++ {
			if i == cnt {
				i = 0
			}
			switch lr[i] {
			case 'L':
				cur = cur.L
			case 'R':
				cur = cur.R
			default:
				panic(1)
			}
			steps++

			if cur.Name[2] == 'Z' {
				stepArray = append(stepArray, steps)
				if _, ok := cycles[State{cur, i}]; !ok {
					fmt.Println("Node:", j, "Cycle:", i, "Steps:", steps, "Factor:", steps/(i+1), "Mod:", steps%(i+1))
					cycles[State{cur, i}] = steps
					part2 *= uint64(steps / (i + 1))
					steps = 0
				} else {
					cycles[State{cur, -1}] = steps
					end = true
				}
			}
		}
	}

	fmt.Println("Part 2:", part2*269) // 269 is the number of "steps" in a cycle
}

func part1() {

	cur := nodes["AAA"] // root node
	// fmt.Println(cur)

	steps := 0
	for i := 0; true; i++ {
		if i == cnt {
			i = 0
		}

		if cur.Name == "ZZZ" {
			break
		}

		switch lr[i] {
		case 'L':
			cur = cur.L
		case 'R':
			cur = cur.R
		default:
			panic(1)
		}

		steps++
	}

	fmt.Println("Part 1: ", steps)
}

type State struct {
	N   *Node
	Pos int
}

type Graph struct {
	Visited    map[State]int
	Jump       []int
	CycleStart int
	Steps      int
}

func BuildGraph(n *Node, lr []rune) Graph {
	g := Graph{
		Visited: make(map[State]int),
		Jump:    make([]int, 0),
	}

	m := len(lr)
	le := 0

	for i := 0; true; i++ {
		i = i % m
		n.Move(lr[i])
		g.Steps++
		if n.Name[2] == 'Z' {
			s := State{n, i}
			if _, ok := g.Visited[s]; !ok {
				g.Visited[s] = g.Steps
				g.Jump = append(g.Jump, g.Steps-le)
				le = g.Steps
			} else {

			}
		}
	}

	return g
}

func (n *Node) Move(d rune) *Node {
	switch d {
	case 'L':
		n = n.L
	case 'R':
		n = n.R
	default:
		panic(1)
	}
	return n
}
