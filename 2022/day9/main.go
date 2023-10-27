package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct{ X, Y int }
type Rope struct{ Knots []Knot }
type Knot struct {
	Location, PreviousLocation Point
	Prev, Next                 *Knot
}

var TailVisited map[Point]bool

func NewRope(size int) Rope {
	b := Rope{make([]Knot, size)}
	var prev *Knot

	prev = nil
	for i := 0; i < size; i++ {
		k := NewKnot()
		k.Prev = prev
		b.Knots[i] = k
		if prev != nil {
			prev.Next = &b.Knots[i]
		}

		prev = &b.Knots[i]
	}
	return b
}

func NewKnot() Knot {
	return Knot{Point{0, 0}, Point{0, 0}, nil, nil}
}

func (r *Knot) MoveHead(direction string, count int) {
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
		r.PreviousLocation.X = r.Location.X
		r.PreviousLocation.Y = r.Location.Y
		r.Location.X += dx
		r.Location.Y += dy
		// fmt.Println(r.Location, r.Next)

		if r.Next != nil {
			r.Next.PullKnot()
		}

	}

}

func (r *Knot) Distance() (*Point, error) {
	if r.Prev != nil {
		return &Point{r.Prev.Location.X - r.Location.X, r.Prev.Location.Y - r.Location.Y}, nil
	}
	return nil, errors.New("at head")
}

func (r *Knot) PullKnot() {
	d, _ := r.Distance()
	// fmt.Printf("%v: %v\n", r, d)
	// var dx, dy int
	// fmt.Println("dist: ", d)

	/* // Doesn't work....
	if d.X > 1 || d.X < -1 || d.Y > 1 || d.Y < -1 {
		r.PreviousLocation.X = r.Location.X
		r.PreviousLocation.Y = r.Location.Y
		r.Location.X = r.Prev.PreviousLocation.X
		r.Location.Y = r.Prev.PreviousLocation.Y
	} else {
		return
	} */
	var dx, dy int

	if d.X == 2 || d.X == -2 {
		dx = d.X / 2
		if d.Y > 0 {
			dy = 1
		} else if d.Y < 0 {
			dy = -1
		}
	} else if d.Y == 2 || d.Y == -2 {
		dy = d.Y / 2
		if d.X > 0 {
			dx = 1
		} else if d.X < 0 {
			dx = -1
		}
	}

	r.Location.X += dx
	r.Location.Y += dy

	if r.Next == nil {
		_, check := TailVisited[Point{r.Location.X, r.Location.Y}]
		if !check {
			fmt.Println("New tail: ", r.Location)
			TailVisited[Point{r.Location.X, r.Location.Y}] = true
		}
	} else {
		r.Next.PullKnot()
	}
}

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	rope := NewRope(10)
	// for i := range rope.Knots {
	// 	fmt.Printf("%v: %p\n", rope.Knots[i], &rope.Knots[i])
	// }
	// return

	TailVisited = make(map[Point]bool, 0)
	TailVisited[Point{0, 0}] = true

	for fileScanner.Scan() {
		line = fileScanner.Text()
		direction, count, _ := strings.Cut(line, " ")
		c, _ := strconv.ParseInt(count, 0, 0)

		fmt.Println(direction, count)
		rope.Knots[0].MoveHead(direction, int(c))
		// fmt.Println(rope.Head, rope.Tail)
		// fmt.Println(rope.Knots)
		for i := range rope.Knots {
			fmt.Printf("%v: %v , %v\n", i, rope.Knots[i].Location, rope.Knots[i].PreviousLocation)
		}
	}
	fmt.Println(rope)
	fmt.Println(len(TailVisited))
}
