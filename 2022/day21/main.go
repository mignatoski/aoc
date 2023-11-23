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
		return Add(h1, h2)
	case '-':
		return Sub(h1, h2)
	case '*':
		return Mult(h1, h2)
	case '/':
		return Div(h1, h2)
	}
	panic("uh oh")
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int) (lcm int, ax int, bx int) {
	result := a * b / GCD(a, b)

	return result, result / a, result / b
}

func Div(h1, h2 Human) Human {
	var result Human

	if h1.CoefN == 0 && h2.CoefN == 0 {
		result.CoefN = 0
		result.CoefD = 1
	} else if h1.CoefN == 0 {
		result.CoefN = h2.CoefN * h1.ConstD
		result.CoefD = h2.CoefD * h1.ConstN
	} else if h2.CoefN == 0 {
		result.CoefN = h1.CoefN * h2.ConstD
		result.CoefD = h1.CoefD * h2.ConstN
	} else {
		panic("too many humans")
	}

	if h1.ConstN == 0 && h2.ConstN == 0 {
		result.ConstN = 0
		result.ConstD = 1
	} else if h1.ConstN == 0 {
		result.ConstN = h2.ConstN
		result.ConstD = h2.ConstD
	} else if h2.ConstN == 0 {
		result.ConstN = h1.ConstN
		result.ConstD = h1.ConstD
	} else {
		result.ConstN = h1.ConstN * h2.ConstD
		result.ConstD = h1.ConstD * h2.ConstN
	}

	gcd := GCD(result.CoefN, result.CoefD)
	result.CoefN /= gcd
	result.CoefD /= gcd
	gcd = GCD(result.ConstN, result.ConstD)
	result.ConstN /= gcd
	result.ConstD /= gcd

	return result
}

func Mult(h1, h2 Human) Human {
	var result Human

	if h1.CoefN == 0 && h2.CoefN == 0 {
		result.CoefN = 0
		result.CoefD = 1
	} else if h1.CoefN == 0 {
		result.CoefN = h2.CoefN * h1.ConstN
		result.CoefD = h2.CoefD * h1.ConstD
	} else if h2.CoefN == 0 {
		result.CoefN = h1.CoefN * h2.ConstN
		result.CoefD = h1.CoefD * h2.ConstD
	} else {
		panic("too many humans")
	}

	if h1.ConstN == 0 && h2.ConstN == 0 {
		result.ConstN = 0
		result.ConstD = 1
	} else if h1.ConstN == 0 {
		result.ConstN = h2.ConstN
		result.ConstD = h2.ConstD
	} else if h2.ConstN == 0 {
		result.ConstN = h1.ConstN
		result.ConstD = h1.ConstD
	} else {
		result.ConstN = h1.ConstN * h2.ConstN
		result.ConstD = h1.ConstD * h2.ConstD
	}

	gcd := GCD(result.CoefN, result.CoefD)
	result.CoefN /= gcd
	result.CoefD /= gcd
	gcd = GCD(result.ConstN, result.ConstD)
	result.ConstN /= gcd
	result.ConstD /= gcd

	return result
}

func Sub(h1, h2 Human) Human {
	var result Human

	result.CoefN = h1.CoefN - h2.CoefN
	result.CoefD = h1.CoefD * h2.CoefD
	if h1.ConstN == 0 && h2.ConstN == 0 {
		result.ConstN = 0
		result.ConstD = 1
	} else if h1.ConstN == 0 {
		result.ConstN = 0 - h2.ConstN
		result.ConstD = h2.ConstD
	} else if h2.ConstN == 0 {
		result.ConstN = h1.ConstN
		result.ConstD = h1.ConstD
	} else {
		lcm, ax, bx := LCM(h1.ConstD, h2.ConstD)
		result.ConstN = (h1.ConstN * ax) - (h2.ConstN * bx)
		result.ConstD = lcm
	}

	gcd := GCD(result.CoefN, result.CoefD)
	result.CoefN /= gcd
	result.CoefD /= gcd
	gcd = GCD(result.ConstN, result.ConstD)
	result.ConstN /= gcd
	result.ConstD /= gcd

	return result
}

func Add(h1, h2 Human) Human {
	var result Human

	result.CoefN = h1.CoefN + h2.CoefN
	result.CoefD = h1.CoefD * h2.CoefD
	if h1.ConstN == 0 && h2.ConstN == 0 {
		result.ConstN = 0
		result.ConstD = 1
	} else if h1.ConstN == 0 {
		result.ConstN = h2.ConstN
		result.ConstD = h2.ConstD
	} else if h2.ConstN == 0 {
		result.ConstN = h1.ConstN
		result.ConstD = h1.ConstD
	} else {
		lcm, ax, bx := LCM(h1.ConstD, h2.ConstD)
		result.ConstN = (h1.ConstN * ax) + (h2.ConstN * bx)
		result.ConstD = lcm
	}

	gcd := GCD(result.CoefN, result.CoefD)
	result.CoefN /= gcd
	result.CoefD /= gcd
	gcd = GCD(result.ConstN, result.ConstD)
	result.ConstN /= gcd
	result.ConstD /= gcd

	return result
}

func Solve(h1, h2 Human) int {
	var result int
	if h1.CoefN == 0 {
		t := h1
		h1 = h2
		h2 = t
	}
	rN := ((h2.ConstN * h1.CoefD * h1.ConstD) - (h1.ConstN * h1.CoefD)) * (h1.CoefD * h1.ConstD)
	rD := (h2.ConstD * h1.CoefD * h1.ConstD) * h1.CoefN * h1.ConstD
	result = rN / rD

	return result
}

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

	fmt.Println(h1, h2)
	fmt.Println("Part 2: ", Solve(h1, h2))

}
