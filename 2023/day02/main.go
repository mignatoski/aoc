package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Game struct {
	Id         int
	IsPossible bool
	Rounds     []Round
	MinCube    Cube
}

type Round struct {
	Cubes Cube
}

type Cube map[string]int

var PART1_LIMIT = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func NewGame() Game {
	return Game{
		Rounds:     make([]Round, 0),
		MinCube:    make(Cube),
		IsPossible: true,
	}
}

func NewRound() Round {
	return Round{
		Cubes: make(Cube),
	}
}

func (c Cube) Power() int {
	p := 1

	for _, v := range c {
		p *= v
	}

	return p
}

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	possible := 0
	power := 0
	for fileScanner.Scan() {
		line = fileScanner.Text()
		game := NewGame()

		g, d, _ := strings.Cut(line, ": ")
		fmt.Sscanf(g, "Game %d", &game.Id)

		rounds := strings.Split(d, "; ")
		for _, r := range rounds {
			round := NewRound()
			c := strings.Split(r, ", ")
			for _, count := range c {
				num, color, _ := strings.Cut(count, " ")
				var numInt int
				fmt.Sscanf(num, "%d", &numInt)
				round.Cubes[color] = numInt
				if numInt > PART1_LIMIT[color] {
					game.IsPossible = false
				}
				if game.MinCube[color] < numInt {
					game.MinCube[color] = numInt
				}
			}
			game.Rounds = append(game.Rounds, round)
		}

		if game.IsPossible {
			possible += game.Id
		}

		power += game.MinCube.Power()

		fmt.Println(game)

	}
	fmt.Println("Part 1: ", possible)

	fmt.Println("Part 2: ", power)
}
