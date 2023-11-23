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
	H         Human
}

type Human struct {
	CoefN, ConstN, CoefD, ConstD int
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

func (m *Monkey) GetHuman() Human {
	if m.Name == "humn" {
		return Human{1, 0, 1, 1}
	}

	if m.Type {
		return m.H
	}

	m1, m2 := FindMonkey(m.M1, monkeys2), FindMonkey(m.M2, monkeys2)
	fmt.Println(m1, m2)
	h1, h2 := m1.GetHuman(), m2.GetHuman()
	fmt.Println(h1, h2, string(m.Op))

	switch m.Op {
	case '+':
		return Human{h1.CoefN + h2.CoefN, (h1.ConstN * h2.ConstD) + (h2.ConstN * h1.ConstD), h1.CoefD * h2.CoefD, h1.ConstD * h2.ConstD}
	case '-':
		return Human{h1.CoefN - h2.CoefN, (h1.ConstN * h2.ConstD) - (h2.ConstN * h1.ConstD), h1.CoefD * h2.CoefD, h1.ConstD * h2.ConstD}
	case '*':
		if h2.CoefN == 0 {
			h2.CoefN = h2.ConstN
		} else if h1.CoefN == 0 {
			h1.CoefN = h1.ConstN
		}
		return Human{h1.CoefN * h2.CoefN, h1.ConstN * h2.ConstN, h1.CoefD * h2.CoefD, h1.ConstD * h2.ConstD}
	case '/':
		if h2.CoefN == 0 {
			h2.CoefN = h2.ConstN
		} else if h1.CoefN == 0 {
			h1.CoefN = h1.ConstN
		}
		return Human{h1.CoefN * h2.CoefD, h1.ConstN * h2.ConstD, h1.CoefD * h2.CoefN, h1.ConstD * h2.ConstN}
	}
	panic("uh oh")
}
func init() {
	monkeys = make([]Monkey, 0)
	monkeys2 = make([]Monkey, 0)
}

func main() {
	inputFile, _ := os.Open("sample.txt")
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
			m.H = Human{0, m.Number, 1, 1}
		}

		monkeys = append(monkeys, m)
		m2 := m
		monkeys2 = append(monkeys2, m2)
	}
	fmt.Println(monkeys)

	root := FindMonkey("root", monkeys)

	fmt.Println("Part 1: ", root.GetNumber())

	root2 := FindMonkey("root", monkeys2)
	m1, m2 := FindMonkey(root2.M1, monkeys2), FindMonkey(root.M2, monkeys2)
	h1, h2 := m1.GetHuman(), m2.GetHuman()

	fmt.Println(h1, h2, m2)
	fmt.Println("Part 2: ", ((h2.ConstN*h1.ConstD)-(h1.ConstN*h2.ConstD))/(h1.ConstD*h2.ConstD))

}
