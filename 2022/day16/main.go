package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"sort"

	"os"
	"strings"

	_ "net/http/pprof"
)

type Valve struct {
	Name            string
	FlowRate        int
	TunnelList      []string
	ConnectedValves []*Valve
	DistanceToValve map[*Valve]int
}

const (
	MINUTES_AVAILABLE int = 30
	TRAVEL_COST       int = 1
	OPEN_COST         int = 1
)

var minutesLeft int
var currentValve *Valve

func CalcPressureAndTimeLeft(c, n *Valve, t int) (int, int) {
	minutesUsed := OPEN_COST + (c.DistanceToValve[n] * TRAVEL_COST)
	effectiveMinutes := t - minutesUsed
	pressureReleased := effectiveMinutes * n.FlowRate
	if effectiveMinutes < 0 {
		effectiveMinutes = 0
	}
	if pressureReleased < 0 {
		pressureReleased = 0
	}

	// fmt.Println(pressureReleased, effectiveMinutes, n.FlowRate, minutesUsed)
	return pressureReleased, effectiveMinutes
}

func RecursiveValves2(remainingValves []*Valve, cv *Valve, next, t, pr int) int {
	// copy map to prevent affecting other recursions
	otherValvesRemaining := make([]*Valve, len(remainingValves))
	copy(otherValvesRemaining, remainingValves)
	nv := otherValvesRemaining[next]
	otherValvesRemaining[next] = otherValvesRemaining[len(otherValvesRemaining)-1]
	otherValvesRemaining = otherValvesRemaining[:len(otherValvesRemaining)-1]

	pressureReleased := 0
	var rt int

	pressureReleased, rt = CalcPressureAndTimeLeft(cv, nv, t)

	if pressureReleased <= 0 {
		return 0
	}

	var maxPressureReleased, npr1, npr2 int

	for nextValve := range otherValvesRemaining {
		npr1 = RecursiveValves2(otherValvesRemaining, nv, nextValve, rt, pressureReleased)
		maxPressureReleased = max(maxPressureReleased, npr1, npr2)
	}
	// fmt.Println(pressureReleased, maxPressureReleased, len(remainingValves))

	return pressureReleased + maxPressureReleased
}

func (v *Valve) RecursiveValves(remainingValves map[*Valve]int, remainingTime, pressureReleased int) int {
	// copy map to prevent affecting other recursions
	otherValvesRemaining := make(map[*Valve]int, 0)
	for k, v := range remainingValves {
		otherValvesRemaining[k] = v
	}
	delete(otherValvesRemaining, v) // remove self from map
	var maxPressureReleased int

	for nextValve := range otherValvesRemaining {
		potentialMinutesUsed := OPEN_COST + (v.DistanceToValve[nextValve] * TRAVEL_COST)
		potentialEffectiveRelease := (remainingTime - potentialMinutesUsed) * nextValve.FlowRate
		if potentialEffectiveRelease > 0 {
			maxPressureReleased = max(maxPressureReleased, nextValve.RecursiveValves(otherValvesRemaining, remainingTime-potentialMinutesUsed, potentialEffectiveRelease))
		}
	}

	return pressureReleased + maxPressureReleased
}

var maxRelease map[string]int

func ValvesString(arr []*Valve) string {
	str := make([]string, len(arr))
	for i, v := range arr {
		str[i] = v.Name
	}
	sort.Slice(str, func(i, j int) bool {
		return str[i] < str[j]
	})

	return strings.Join(str, ",")
}

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	// Load Valves

	valves := make(map[string]*Valve, 0)
	var line string
	for fileScanner.Scan() {
		valve := Valve{}
		valve.ConnectedValves = make([]*Valve, 0)
		line = fileScanner.Text()
		fmt.Sscanf(line, "Valve %s has flow rate=%d;", &valve.Name, &valve.FlowRate)
		_, tunnels, _ := strings.Cut(strings.Replace(line, "valves", "valve", 2), "valve ")

		valve.TunnelList = strings.Split(tunnels, ", ")
		valves[valve.Name] = &valve
		valve.DistanceToValve = make(map[*Valve]int, 0)

		// fmt.Println(valve)
	}

	for _, v := range valves {
		for _, cv := range v.TunnelList {
			v.ConnectedValves = append(v.ConnectedValves, valves[cv])
		}
		// fmt.Println(v.Name, v.FlowRate, v.ConnectedValves)
	}

	// Determine distance to each valve

	for _, v := range valves {
		valvesToVisit := make(map[*Valve]int, 0)
		valvesToVisit[v] = 0

		for len(valvesToVisit) > 0 {
			keys := make([]*Valve, len(valvesToVisit))

			i := 0
			for k := range valvesToVisit {
				keys[i] = k
				i++
			}
			for _, valve := range keys {
				v.DistanceToValve[valve] = valvesToVisit[valve]
				for _, nextValve := range valve.ConnectedValves {
					if _, exists := v.DistanceToValve[nextValve]; !exists {
						valvesToVisit[nextValve] = v.DistanceToValve[valve] + 1
					}
				}
				delete(valvesToVisit, valve)
			}
		}
		// fmt.Println(v.Name, v.FlowRate, v.ConnectedValves, len(v.DistanceToValve))
	}

	// Part 2 - With partner
	currentValve = valves["AA"]
	minutesLeft = MINUTES_AVAILABLE - 4 // teach elephant

	positiveValves := make([]*Valve, 0)
	for _, v := range valves {
		if v.FlowRate > 0 {
			positiveValves = append(positiveValves, v)
		}
	}

	sort.Slice(positiveValves, func(i, j int) bool {
		return positiveValves[i].Name < positiveValves[j].Name
	})

	// for _, v := range positiveValves {
	// 	println(v.Name, v.FlowRate)
	// 	for k, d := range v.DistanceToValve {
	// 		println("-- ", k.Name, d)
	// 	}
	// }

	maxRelease = make(map[string]int)
	PopulateValveMap(Queue{
		openedValves:   make([]*Valve, 0),
		current:        currentValve,
		timeLeft:       minutesLeft,
		pressureRelief: 0,
	}, positiveValves)

	// for k := range positiveValves {
	// 	npr1 = RecursiveValves2(positiveValves, currentValve, currentValve, k, minutesLeft, minutesLeft)
	// 	// fmt.Println(count, npr1, npr2)
	// 	potentialMaxPressureRelease = max(potentialMaxPressureRelease, npr1, npr2)
	// }

	var maxPr int
	for _, v := range maxRelease {
		maxPr = max(maxPr, v)
	}
	fmt.Println("Part 1:", maxPr)

	maxPr = 0
	for k1, p1 := range maxRelease {
		for k2, p2 := range maxRelease {
			found := false
			a1, a2 := strings.Split(k1, ","), strings.Split(k2, ",")
			for _, v1 := range a1 {
				for _, v2 := range a2 {
					if v1 == v2 {
						found = true
						break
					}
				}
			}
			if !found {
				maxPr = max(maxPr, p1+p2)
			}
		}
	}
	fmt.Println("Part 2:", maxPr)

	// fmt.Println("Part 2 Pressure Released: ", potentialMaxPressureRelease)

	// Graph Traversal - Part 1

	// currentValve = valves["AA"]
	// minutesLeft = MINUTES_AVAILABLE

	// positiveValves := make(map[*Valve]int, 0)
	// for _, v := range valves {
	// positiveValves[v] = 0
	// }
	// var potentialMaxPressureRelease intjk
	// for k := range positiveValves {
	// potentialMinutesUsed := OPEN_COST + (currentValve.DistanceToValve[k] * TRAVEL_COST)
	// potentialEffectiveRelease := (minutesLeft - potentialMinutesUsed) * k.FlowRate
	// potentialMaxPressureRelease = max(potentialMaxPressureRelease, k.RecursiveValves(positiveValves, minutesLeft-potentialMinutesUsed, potentialEffectiveRelease))
	// }

	// fmt.Println("Part 1 Pressure Released: ", potentialMaxPressureRelease)

}

type Queue struct {
	openedValves             []*Valve
	current                  *Valve
	timeLeft, pressureRelief int
}

func PopulateValveMap(iq Queue, positiveValves []*Valve) {

	aq := make([]*Queue, 1)
	aq[0] = &iq

	for len(aq) > 0 {
		q := aq[0]
		for _, v := range positiveValves {
			found := false
			for _, ov := range q.openedValves {
				if v.Name == ov.Name {
					found = true
				}
			}
			if !found && q.current.DistanceToValve[v] <= q.timeLeft {
				newOpenedValves := make([]*Valve, len(q.openedValves))
				copy(newOpenedValves, q.openedValves)
				newOpenedValves = append(newOpenedValves, v)
				key := ValvesString(newOpenedValves)
				pr, tl := CalcPressureAndTimeLeft(q.current, v, q.timeLeft)
				if tl > 0 {
					nq := &Queue{
						openedValves:   newOpenedValves,
						current:        v,
						timeLeft:       tl,
						pressureRelief: pr + q.pressureRelief,
					}
					aq = append(aq, nq)
				}
				// fmt.Println(key, pr+q.pressureRelief, pr, q.pressureRelief)
				maxRelease[key] = max(maxRelease[key], pr+q.pressureRelief)

			}
		}
		aq = aq[1:]
	}

}
