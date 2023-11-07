package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Point struct {
	X, Y int
}

type Section struct {
	X1, X2 int
}

func (s *Section) Length() int {
	return 1 + s.X2 - s.X1
}

type Sensor struct {
	Location, NearestBeacon Point
}

type Line struct {
	a, b, c int
}

func (p *Point) Coef(slope int) Line {
	// y - y1 = m(x - x1)
	// 0 = mx + (-1)y + (y1 - mx1)

	return Line{slope, -1, (p.Y - (slope * p.X))}
}

func (l Line) Intersect(l2 Line) (p Point) {
	p.X = (l.b*l2.c - l2.b*l.c) / (l.a*l2.b - l2.a*l.b)
	p.Y = (l2.a*l.c - l.a*l2.c) / (l.a*l2.b - l2.a*l.b)

	return
}

func (s *Sensor) BoundaryIntersectPoints(s2 *Sensor) []Point {
	points := make([]Point, 4)
	boundaries := make([]Point, 0)

	points[0] = Point{s.Location.X - s.DistanceToBeacon() - 1, s.Location.Y}
	points[1] = Point{s.Location.X + s.DistanceToBeacon() + 1, s.Location.Y}
	points[2] = Point{s2.Location.X - s2.DistanceToBeacon() - 1, s2.Location.Y}
	points[3] = Point{s2.Location.X + s2.DistanceToBeacon() + 1, s2.Location.Y}

	for i := 0; i < 2; i++ {
		for j := 2; j < 4; j++ {
			bp1 := points[i].Coef(1).Intersect(points[j].Coef(-1))
			bp2 := points[i].Coef(-1).Intersect(points[j].Coef(1))
			if bp1.X >= 0 && bp1.X <= 4_000_000 && bp1.Y >= 0 && bp1.Y <= 4_000_000 {
				boundaries = append(boundaries, bp1)
			}
			if bp2.X >= 0 && bp2.X <= 4_000_000 && bp2.Y >= 0 && bp2.Y <= 4_000_000 {
				boundaries = append(boundaries, bp2)
			}
		}
	}

	return boundaries
}

func (s *Sensor) DistanceToBeacon() int {
	return s.DistanceToPoint(s.NearestBeacon)
}

func (s *Sensor) DistanceToPoint(p Point) int {
	return int(math.Abs(float64(s.Location.X-p.X)) + math.Abs(float64(s.Location.Y-p.Y)))
}

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	sensors := make([]Sensor, 0)
	beacons := make(map[Point]bool, 0)
	count := 0
	for fileScanner.Scan() {
		// if count > 5 {
		// break
		// }
		line = fileScanner.Text()
		sensor := Sensor{}

		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.Location.X, &sensor.Location.Y, &sensor.NearestBeacon.X, &sensor.NearestBeacon.Y)
		sensors = append(sensors, sensor)
		beacons[Point{sensor.NearestBeacon.X, sensor.NearestBeacon.Y}] = true
		count++
		// fmt.Println(int(math.Abs(float64(sensor.Location.Y-2_000_000))) - sensor.DistanceToBeacon())
	}

	//          B
	//          #
	//          #
	//          #
	//			S
	// 2M    #######

	// offset = 1
	// dtb = 4
	// xr = 3
	// minX = X - 3
	// maxX = X + 3

	sections := make([]Section, 0)
	for _, s := range sensors {
		offset := 2_000_000 - s.Location.Y
		xRange := s.DistanceToBeacon() - int(math.Abs(float64(offset)))
		if xRange >= 0 {
			minX := s.Location.X - xRange
			maxX := s.Location.X + xRange

			sections = append(sections, Section{minX, maxX})
		}

	}
	sort.Slice(sections, func(i, j int) bool {
		return sections[i].X1 < sections[j].X1
	})
	fmt.Println(sections)

	tmp := sections[0]
	condensed := make([]Section, 0)
	for _, v := range sections {
		if v.X1 > tmp.X2 {
			condensed = append(condensed, tmp)
			tmp = v
		}
		if v.X2 > tmp.X2 {
			tmp.X2 = v.X2
		}
	}
	condensed = append(condensed, tmp)
	fmt.Println(condensed)

	var nonBeaconCount int
	for _, v := range condensed {
		nonBeaconCount += v.Length()
	}

	fmt.Println(nonBeaconCount)
	for k := range beacons {
		if k.Y == 2_000_000 {
			for _, v := range condensed {
				if k.X >= v.X1 && k.X <= v.X2 {
					nonBeaconCount--
				}
			}
		}
	}
	fmt.Println("Part 1: ", nonBeaconCount)

	// Part 2
	// dist between two sensors and total of their beacon dist + 1 are equal
	// calc distance betwwen all other sensors

	possibleBeacons := make([]Point, 0)
	for i, s := range sensors {
		for j := i + 1; j < len(sensors); j++ {
			s2 := sensors[j]
			mustBeOneOrLess := s.DistanceToPoint(s2.Location) - s.DistanceToBeacon() - s2.DistanceToBeacon()
			if mustBeOneOrLess <= 1 {
				possibleBeacons = append(possibleBeacons, s.BoundaryIntersectPoints(&s2)...)
			}
		}

	}
	// fmt.Println(possibleBeacons)

	for _, b := range possibleBeacons {
		skip := false
		for _, s := range sensors {
			if s.DistanceToPoint(b) <= s.DistanceToBeacon() {
				skip = true
				break
			}
		}
		if !skip {
			fmt.Println("Part 2:")
			fmt.Println(b)
			fmt.Println("Tuning Freq: ", b.X*4000000+b.Y)
			break
		}
	}

}
