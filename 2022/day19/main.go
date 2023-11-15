package main

import (
	"bufio"
	"fmt"
	"os"
)

type Blueprint struct {
	Id        int
	RobotCost []Cost
}

type Cost struct {
	Ore, Clay, Obsidian int
}

type State struct {
	Minute            int
	Robots, Resources []int
	BP                Blueprint
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
		b := Blueprint{RobotCost: make([]Cost, 4)}
		line = fileScanner.Text()
		fmt.Sscanf(line,
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&b.Id,
			&b.RobotCost[ORE].Ore,
			&b.RobotCost[CLAY].Ore,
			&b.RobotCost[OBSIDIAN].Ore, &b.RobotCost[OBSIDIAN].Clay,
			&b.RobotCost[GEODE].Ore, &b.RobotCost[GEODE].Obsidian)
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
	s := State{
		Minute:    1,
		Robots:    make([]int, 4),
		Resources: make([]int, 4),
		BP:        b,
	}
	return b.Id * s.MaxGeodes()
}

func (s State) MaxGeodes() int {
	if s.Minute == MINUTES {
		s.Collect()
		return s.Resources[GEODE]
	}

	s.Minute++
	var maxGeode int
	maxGeode = max(maxGeode, s.Collect().MaxGeodes())
	for i := 0; i < 4; i++ {
		if s.CanBuy(Resource(i)) {
			maxGeode = max(maxGeode, s.Collect().Buy(Resource(i)).MaxGeodes())
		}
	}

	return maxGeode
}

func (s State) CanBuy(r Resource) bool {
	if s.Resources[ORE] >= s.BP.RobotCost[r].Ore && s.Resources[CLAY] >= s.BP.RobotCost[r].Clay && s.Resources[OBSIDIAN] >= s.BP.RobotCost[r].Obsidian {
		return true
	}
	return false
}

func (s State) Buy(r Resource) State {
	s.Robots[r]++
	s.Resources[ORE] -= s.BP.RobotCost[r].Ore
	s.Resources[CLAY] -= s.BP.RobotCost[r].Clay
	s.Resources[OBSIDIAN] -= s.BP.RobotCost[r].Obsidian
	return s

}

func (s State) Collect() State {
	for i := 0; i < 4; i++ {
		s.Resources[i] += s.Robots[i]
	}
	return s
}
