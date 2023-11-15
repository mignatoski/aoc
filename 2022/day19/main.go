package main

import (
	"bufio"
	"fmt"
	"os"
)

type Blueprint struct {
	Id                         int
	Ore, Clay, Obsidian, Geode Cost
}

type Cost struct {
	Ore, Clay, Obsidian int
}

type State struct {
	Minute            int
	Robots, Resources []int
}

type Resource int

const (
	ORE Resource = iota
	CLAY
	OBSIDIAN
	GEODE
)

const (
	MINUTES = 24
)

var blueprints []Blueprint

func main() {
	inputFile, _ := os.Open("sample.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	for fileScanner.Scan() {
		b := Blueprint{}
		line = fileScanner.Text()
		fmt.Sscanf(line,
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&b.Id, &b.Ore.Ore, &b.Clay.Ore, &b.Obsidian.Ore, &b.Obsidian.Clay, &b.Geode.Ore, &b.Geode.Obsidian)
		blueprints = append(blueprints, b)
	}
	fmt.Println(blueprints)

	totalQualityLevel := 0
	for _, b := range blueprints {
		totalQualityLevel += b.QualityLevel()
	}

	fmt.Println("Part 1:", totalQualityLevel)
}

func (b Blueprint) QualityLevel() int {

	return 0
}

func (s State) MaxGeodes(b Blueprint) {

}
