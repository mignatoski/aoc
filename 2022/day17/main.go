package main

import (
	"bufio"
	"fmt"
	"os"
)

var jetPos int
var jets []rune
var numRocks int
var rockPattern [][]Point

func nextJet() (next rune) {
	next = jets[jetPos]
	jetPos = (jetPos + 1) % len(jets)
	return
}

type Rock struct {
	Points []Point
}

type Point struct {
	X, Y int
}

func init() {
	rockPattern = [][]Point{
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}},         // -
		{{0, 1}, {1, 0}, {1, 1}, {1, 2}, {2, 1}}, // +
		{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}, // L backwards
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}},         // |
		{{0, 0}, {1, 0}, {0, 1}, {1, 1}},         // Square
	}
}

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	fileScanner.Scan()
	line = fileScanner.Text()
	jets = []rune(line)

	var rock *Rock
	for numRocks < 2023 {
		// Rock Falls
		rock = RockFalls(rock)

		// Rock Push by Jet
		rock.PushByJet()
	}

	fmt.Println(jets[0])
}

func (rock *Rock) PushByJet() {
	jet := nextJet()
	switch jet {
	case '<':
	case '>':
	}

}

func RockFalls(rock *Rock) *Rock {
	numRocks++

	return &Rock{}
}
