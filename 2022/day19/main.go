package main

import (
	"bufio"
	"fmt"
	"os"
)

type Blueprint struct {
	Id        int
	RobotCost []Cost
	MaxGeodes int
	Stack     []*State
	Most      []int
}

type Cost struct {
	Ore, Clay, Obsidian int
}

type State struct {
	Minute            int
	Robots, Resources []int
	BP                *Blueprint
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
		cost := b.RobotCost
		b.Most = make([]int, 4)
		b.Most[ORE] = max(cost[ORE].Ore, cost[CLAY].Ore, cost[OBSIDIAN].Ore, cost[GEODE].Ore)
		b.Most[CLAY] = cost[OBSIDIAN].Clay
		b.Most[OBSIDIAN] = cost[GEODE].Obsidian
		b.Most[GEODE] = 100_000
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
		BP:        &b,
	}
	s.Robots[ORE] = 1
	b.Stack = make([]*State, 1)
	b.Stack[0] = &s

	b.UpdateMaxGeodes()

	fmt.Println(b.Id, b.MaxGeodes)
	return b.Id * b.MaxGeodes
}

func (b *Blueprint) UpdateMaxGeodes() {
	s := b.Stack[0].Copy()
	b.Stack = b.Stack[1:]

	if s.Minute == MINUTES {
		s.Collect()
		s.BP.MaxGeodes = max(s.BP.MaxGeodes, s.Resources[GEODE])
	}

	if s.PotentialGeodes() > s.BP.MaxGeodes {
		for i := 3; i >= 0; i-- {
			ns2 := s.Copy()
			if ns2.CanBuy(Resource(i)) {
				ns2.Collect().Buy(Resource(i))
				ns2.Minute++
				b.Stack = append(b.Stack, &ns2)
				b.UpdateMaxGeodes()
			}
		}
		ns := s.Copy()
		ns.Collect()
		ns.Minute++
		b.Stack = append(b.Stack, &ns)
		b.UpdateMaxGeodes()
	}
}
func (s *State) TimeLeft() int {
	return MINUTES - s.Minute
}

func (s *State) PotentialGeodes() int {
	geodeRate := s.Robots[GEODE]
	potentialGeodes := s.Resources[GEODE]
	for i := 0; i < s.TimeLeft(); i++ {
		potentialGeodes = geodeRate
		geodeRate++
	}
	return potentialGeodes
}

func (s *State) CanBuy(r Resource) bool {
	cost := s.BP.RobotCost[r]

	if true &&
		s.Robots[r] < s.BP.Most[r] &&
		s.BP.Most[r]*s.TimeLeft() > s.Resources[r] &&
		s.Resources[ORE] >= cost.Ore &&
		s.Resources[CLAY] >= cost.Clay &&
		s.Resources[OBSIDIAN] >= cost.Obsidian {
		return true
	}
	return false
}

func (s *State) Buy(r Resource) *State {
	s.Robots[r]++
	s.Resources[ORE] -= s.BP.RobotCost[r].Ore
	s.Resources[CLAY] -= s.BP.RobotCost[r].Clay
	s.Resources[OBSIDIAN] -= s.BP.RobotCost[r].Obsidian
	return s

}

func (s *State) Collect() *State {
	for i := 0; i < 4; i++ {
		s.Resources[i] += s.Robots[i]
	}
	return s
}

func (s *State) Copy() State {
	ns := State{
		Robots:    make([]int, 4),
		Resources: make([]int, 4),
	}
	ns.Minute = s.Minute
	copy(ns.Robots, s.Robots)
	copy(ns.Resources, s.Resources)
	ns.BP = s.BP

	return ns
}
