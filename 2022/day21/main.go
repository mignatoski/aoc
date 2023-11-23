package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Monkey struct {
	Name, Job string
	Number    int
	Type      bool
	Op        rune
	M1, M2    string
}

func FindMonkey(name string, l []Monkey) *Monkey {
	for _, m := range l {
		if m.Name == name {
			return &m
		}
	}
	panic("couldn't find monkey")
}

func (m *Monkey) GetNumber() int {
	if m.Type {
		return m.Number
	}

	m1, m2 := FindMonkey(m.M1, monkeys), FindMonkey(m.M2, monkeys)
	n1, n2 := m1.GetNumber(), m2.GetNumber()

	switch m.Op {
	case '+':
		return n1 + n2
	case '-':
		return n1 - n2
	case '*':
		return n1 * n2
	case '/':
		return n1 / n2
	}
	panic("uh oh")
}

var monkeys, monkeys2 []Monkey

func init() {
	monkeys = make([]Monkey, 0)
	monkeys2 = make([]Monkey, 0)
}

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	for fileScanner.Scan() {
		line = fileScanner.Text()
		m := Monkey{}
		m.Name, m.Job, _ = strings.Cut(line, ": ")
		args := strings.Split(m.Job, " ")
		if len(args) == 3 {
			m.Type = false
			m.M1, m.M2 = args[0], args[2]
			m.Op = []rune(args[1])[0]
		} else {
			m.Type = true
			fmt.Sscanf(args[0], "%d", &m.Number)
		}

		monkeys = append(monkeys, m)
		m2 := m
		monkeys2 = append(monkeys2, m2)
	}
	fmt.Println(monkeys)

	root := FindMonkey("root", monkeys)

	fmt.Println("Part 1: ", root.GetNumber())

}
