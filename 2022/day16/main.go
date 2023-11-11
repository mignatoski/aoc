package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	return pressureReleased, effectiveMinutes
}

func RecursiveValves2(remainingValves map[*Valve]*int, cv1, cv2, nv *Valve, t1, t2, move int) int {
	// copy map to prevent affecting other recursions
	otherValvesRemaining := make(map[*Valve]*int, 0)
	var optimisticPressure int
	for k := range remainingValves {
		otherValvesRemaining[k] = &optimisticPressure
		optimisticPressure += max(t1, t2) * k.FlowRate
	}
	delete(otherValvesRemaining, nv) // remove self from map

	pressureReleased := 0
	var rt1, rt2 int

	if move == 1 {
		pressureReleased, rt1 = CalcPressureAndTimeLeft(cv1, nv, t1)
		rt2 = t2
		cv1 = nv
	} else {
		pressureReleased, rt2 = CalcPressureAndTimeLeft(cv2, nv, t2)
		rt1 = t1
		cv2 = nv
	}
	if pressureReleased <= 0 {
		return 0
	}

	var maxPressureReleased, npr1, npr2 int

	for nextValve := range otherValvesRemaining {
		if optimisticPressure > maxPressureReleased {

			if rt1 > 0 {
				npr1 = RecursiveValves2(otherValvesRemaining, cv1, cv2, nextValve, rt1, rt2, 1)
			}
			if rt2 > 0 {
				npr2 = RecursiveValves2(otherValvesRemaining, cv1, cv2, nextValve, rt1, rt2, 2)
			}
			maxPressureReleased = max(maxPressureReleased, npr1, npr2)
		}
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

func main() {
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
		fmt.Println(v.Name, v.FlowRate, v.ConnectedValves, len(v.DistanceToValve))
	}

	// Part 2 - With partner
	currentValve = valves["AA"]
	minutesLeft = MINUTES_AVAILABLE - 4 // teach elephant

	positiveValves := make(map[*Valve]*int, 0)
	op := 99999999
	for _, v := range valves {
		if v.FlowRate > 0 {
			positiveValves[v] = &op
		}
	}
	var potentialMaxPressureRelease, npr1, npr2, count int
	for k := range positiveValves {
		count++
		npr1 = RecursiveValves2(positiveValves, currentValve, currentValve, k, minutesLeft, minutesLeft, 1)
		npr2 = RecursiveValves2(positiveValves, currentValve, currentValve, k, minutesLeft, minutesLeft, 2)
		fmt.Println(count, npr1, npr2)
		potentialMaxPressureRelease = max(potentialMaxPressureRelease, npr1, npr2)
	}

	fmt.Println("Part 2 Pressure Released: ", potentialMaxPressureRelease)

	return

	// Graph Traversal - Part 1

	currentValve = valves["AA"]
	minutesLeft = MINUTES_AVAILABLE

	// positiveValves := make(map[*Valve]int, 0)
	// for _, v := range valves {
	// positiveValves[v] = 0
	// }
	// var potentialMaxPressureRelease int
	// for k := range positiveValves {
	// potentialMinutesUsed := OPEN_COST + (currentValve.DistanceToValve[k] * TRAVEL_COST)
	// potentialEffectiveRelease := (minutesLeft - potentialMinutesUsed) * k.FlowRate
	// potentialMaxPressureRelease = max(potentialMaxPressureRelease, k.RecursiveValves(positiveValves, minutesLeft-potentialMinutesUsed, potentialEffectiveRelease))
	// }

	fmt.Println("Part 1 Pressure Released: ", potentialMaxPressureRelease)

}
