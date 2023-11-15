package main

import (
	"bufio"
	"fmt"
	"os"
)

var cubes []Cube
var bounds Cube

type Cube struct {
	X, Y, Z int
}

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	cubes = make([]Cube, 0)
	for fileScanner.Scan() {
		line = fileScanner.Text()
		c := Cube{}
		fmt.Sscanf(line, "%d,%d,%d", &c.X, &c.Y, &c.Z)
		cubes = append(cubes, c)
		bounds.X = max(bounds.X, c.X)
		bounds.Y = max(bounds.Y, c.Y)
		bounds.Z = max(bounds.Z, c.Z)

	}
	fmt.Println(cubes)

	var sides int
	for _, c := range cubes {
		sides += 6
		for i := -1; i <= 1; i += 2 {
			if inCubes(Cube{c.X + i, c.Y, c.Z}) {
				sides--
			}
			if inCubes(Cube{c.X, c.Y + i, c.Z}) {
				sides--
			}
			if inCubes(Cube{c.X, c.Y, c.Z + i}) {
				sides--
			}
		}
	}

	fmt.Println("Part 1: ", sides)

	sides = 0
	for _, c := range cubes {
		sides += 6
		var t Cube
		for i := -1; i <= 1; i += 2 {
			t = Cube{c.X + i, c.Y, c.Z}
			if inCubes(t) || inAirBubble(t) {
				sides--
			}
			t = Cube{c.X, c.Y + i, c.Z}
			if inCubes(t) || inAirBubble(t) {
				sides--
			}
			t = Cube{c.X, c.Y, c.Z + i}
			if inCubes(t) || inAirBubble(t) {
				sides--
			}
		}
	}

	fmt.Println("Part 2: ", sides, bounds, inAirBubble(Cube{2, 2, 5}))
}

func inAirBubble(c Cube) bool {
	check := make([]Cube, 1)
	visited := make(map[Cube]bool)
	check[0] = c
	visited[c] = true

	var cube, t Cube
	for len(check) > 0 {
		cube = check[0]
		if cube.X == 0 || cube.X > bounds.X {
			return false
		}
		if cube.Y == 0 || cube.Y > bounds.Y {
			return false
		}
		if cube.Z == 0 || cube.Z > bounds.Z {
			return false
		}

		for i := -1; i <= 1; i += 2 {
			t = Cube{cube.X + i, cube.Y, cube.Z}
			if !visited[t] && !inCubes(t) {
				visited[t] = true
				check = append(check, t)
			}
			t = Cube{cube.X, cube.Y + i, cube.Z}
			if !visited[t] && !inCubes(t) {
				visited[t] = true
				check = append(check, t)
			}
			t = Cube{cube.X, cube.Y, cube.Z + i}
			if !visited[t] && !inCubes(t) {
				visited[t] = true
				check = append(check, t)
			}
		}
		check = check[1:]
	}

	return true
}

func inCubes(c Cube) bool {
	for i := 0; i < len(cubes); i++ {
		if c == cubes[i] {
			return true
		}
	}
	return false
}
