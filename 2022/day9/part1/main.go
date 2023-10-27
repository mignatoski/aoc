package part1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct{ X, Y int }
type Rope struct {
	Head, Tail  Point
	TailVisited map[Point]bool
}

func NewRope() Rope {
	v := make(map[Point]bool, 0)
	v[Point{0, 0}] = true
	return Rope{Point{0, 0}, Point{0, 0}, v}
}

func (r *Rope) MoveHead(direction string, count int) {
	dx, dy := 0, 0
	switch direction {
	case "U":
		dy = 1
	case "D":
		dy = -1
	case "R":
		dx = 1
	case "L":
		dx = -1
	}
	for i := 0; i < count; i++ {
		r.Head = Point{r.Head.X + dx, r.Head.Y + dy}
		if r.Distance().TooFar() {
			r.moveTail()
		}
		// fmt.Println(r.Head, r.Tail)
	}

}

func (r *Rope) Distance() Point {
	return Point{r.Head.X - r.Tail.X, r.Head.Y - r.Tail.Y}
}

func (r *Rope) moveTail() {
	d := r.Distance()
	var dx, dy int

	if d.X > 1 {
		dx = 1
		dy = d.Y
	} else if d.X < -1 {
		dx = -1
		dy = d.Y
	} else if d.Y > 1 {
		dx = d.X
		dy = 1
	} else if d.Y < -1 {
		dx = d.X
		dy = -1
	}

	x, y := r.Tail.X+dx, r.Tail.Y+dy

	r.Tail = Point{x, y}
	r.TailVisited[r.Tail] = true
}

func (p Point) TooFar() bool {
	return math.Abs(float64(p.X)) > 1 || math.Abs(float64(p.Y)) > 1
}

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	rope := NewRope()
	for fileScanner.Scan() {
		line = fileScanner.Text()
		direction, count, _ := strings.Cut(line, " ")
		c, _ := strconv.ParseInt(count, 0, 0)

		fmt.Println(direction, count)
		rope.MoveHead(direction, int(c))
		// fmt.Println(rope.Head, rope.Tail)
	}
	fmt.Println(rope)
	fmt.Println(len(rope.TailVisited))
}
