package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Blueprint struct {
	Id        int
	RobotCost []Cost
	MaxGeodes int
	Stack     []*State
	Most      Cost
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
		b.Most.Ore = max(cost[ORE].Ore, cost[CLAY].Ore, cost[OBSIDIAN].Ore, cost[GEODE].Ore)
		b.Most.Clay = cost[OBSIDIAN].Clay
		b.Most.Obsidian = cost[GEODE].Obsidian
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

	var cur *State

	for len(b.Stack) > 0 {
		sort.Slice(b.Stack, func(i, j int) bool {
			return b.Stack[i].Minute > b.Stack[j].Minute
			// return b.Stack[i].PotentialGeodes() > b.Stack[j].PotentialGeodes()
		})
		cur = b.Stack[0]
		cur.MaxGeodes()

		b.Stack = b.Stack[1:]
	}
	fmt.Println(b.Id, b.MaxGeodes)
	return b.Id * b.MaxGeodes
}

func (s State) MaxGeodes() {
	ns := s.Copy()
	if ns.Minute == MINUTES {
		fmt.Println("END: ", ns)
		ns.Collect()
		ns.BP.MaxGeodes = max(ns.BP.MaxGeodes, ns.Resources[GEODE])
	}

	if s.PotentialGeodes() > s.BP.MaxGeodes {
		ns.Minute++
		ns.Collect()
		s.BP.Stack = append(s.BP.Stack, &ns)
		fmt.Println("Nothing: ", ns)
		for i := 3; i >= 0; i-- {
			ns2 := s.Copy()
			ns2.Minute++
			if ns2.CanBuy(Resource(i)) {
				fmt.Println("BUY: ", ns)
				ns2.Collect().Buy(Resource(i))
				s.BP.Stack = append(s.BP.Stack, &ns2)
			}
		}
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

	/* maxOre := s.TimeLeft() * (s.BP.RobotCost[CLAY].Ore +
		s.BP.RobotCost[OBSIDIAN].Ore +
		s.BP.RobotCost[GEODE].Ore)

	maxClay := s.TimeLeft() * (s.BP.RobotCost[OBSIDIAN].Clay +
		s.BP.RobotCost[GEODE].Clay)

	maxObsidian := s.TimeLeft() * (s.BP.RobotCost[GEODE].Obsidian) */

	if true &&
		s.Robots[ORE] < s.BP.Most.Ore &&
		s.Robots[CLAY] < s.BP.Most.Clay &&
		s.Robots[OBSIDIAN] < s.BP.Most.Obsidian &&
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
