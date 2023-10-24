package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line, mode string
	root := NewDir("/", nil)
	c := root
	for fileScanner.Scan() {
		line = fileScanner.Text()
		if strings.HasPrefix(line, "$ ") {
			mode = "cmd"
			cmd, arg, _ := strings.Cut(line[2:], " ")
			switch cmd {
			case "cd":
				switch arg {

				case "/":
					c = root

				case "..":
					c = c.Parent

				default:
					for _, child := range c.Children {
						if child.Name == arg {
							c = child
						}
					}
				}
			case "ls":
				// ignore
			default:
				panic("unknown command")
			}
		} else {
			mode = "output"
			v1, v2, _ := strings.Cut(line, " ")
			if v1 == "dir" {
				c.Children = append(c.Children, NewDir(v2, c))
			} else {
				size, _ := strconv.ParseInt(v1, 0, 0)
				c.Children = append(c.Children, NewFile(v2, int(size), c))
			}
		}
		fmt.Println(mode, line, c.Name)
	}

	// fmt.Println(root)
	root.Print("")

	fmt.Println("part 1: ", root.Answer())

	SpaceAvailable := 70_000_000
	TargetUnused := 30_000_000
	SpaceUsed := root.TotalSize()
	SpaceUnused := SpaceAvailable - SpaceUsed
	ToDelete := TargetUnused - SpaceUnused
	fmt.Println(SpaceAvailable, TargetUnused, SpaceUsed, SpaceUnused, ToDelete)

	fmt.Println("part 2: ", root.Part2(ToDelete, SpaceAvailable))

}

type Tree struct {
	Name, Type string
	Size       int
	Children   []*Tree
	Parent     *Tree
}

func (t *Tree) Part2(toDelete int, smallest int) int {
	if t.Type == "dir" {
		ts := t.TotalSize()
		if ts >= toDelete && ts < smallest {
			smallest = ts
		}
		for _, child := range t.Children {
			smallest = child.Part2(toDelete, smallest)
		}
	}
	return smallest
}

func (t *Tree) Answer() int {
	sum := 0
	if t.Type == "dir" {
		ts := t.TotalSize()
		if ts <= 100000 {
			fmt.Println(t.Name, t.TotalSize())
			sum = ts
		}
		for _, child := range t.Children {
			sum += child.Answer()
		}
	}
	return sum
}

func (t *Tree) TotalSize() int {
	switch t.Type {
	case "file":
		return t.Size
	default:
		sum := 0
		for _, child := range t.Children {
			sum += child.TotalSize()
		}
		return sum
	}
}

func (t *Tree) Print(s string) {
	fmt.Println(s, t.Name, "(", t.Type, t.TotalSize(), ")")
	for _, child := range t.Children {
		child.Print(s + "  ")
	}
}

func NewFile(name string, size int, parent *Tree) *Tree {
	return &Tree{Name: name, Type: "file", Size: size, Parent: parent}
}

func NewDir(name string, parent *Tree) *Tree {
	return &Tree{Name: name, Type: "dir", Parent: parent, Children: make([]*Tree, 0)}
}
