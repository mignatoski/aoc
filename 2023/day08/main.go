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

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string

	fileScanner.Scan()
	lr := []rune(fileScanner.Text())

	fileScanner.Scan() // read blank line

	nodes := make(map[string]*Node, 0)
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
		fmt.Println(l, r, n.L, n.R)

		if n.Name[2] == 'A' {
			curs = append(curs, n)
		}
	}

	fmt.Println(lr, nodes)

	cur := nodes["AAA"] // root node
	fmt.Println(cur)

	cnt := len(lr)
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

	steps = 0
	fmt.Println(len(curs))
	for i := 0; true; i++ {
		if i == cnt {
			i = 0
		}

		end := true
		for j, cur := range curs {
			if cur.Name[2] != 'Z' {
				end = false
			} else {
				fmt.Println("Node:", j, "Cycle:", i)
			}
			switch lr[i] {
			case 'L':
				cur = cur.L
			case 'R':
				cur = cur.R
			default:
				panic(1)
			}
		}
		if end {
			break
		}

		steps++
	}

	fmt.Println("Part 2: ", steps)
}
