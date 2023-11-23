package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Number struct {
	Val, InitialPosition int
}

type List []Number

func (l List) String() string {
	var sb strings.Builder
	var bz strings.Builder
	found := false
	C := len(l) - 1

	for _, n := range l {
		if n.Val == 0 {
			found = true
		}

		if found {
			sb.WriteString(fmt.Sprint(n.Val%C, ", "))
		} else {
			bz.WriteString(fmt.Sprint(n.Val%C, ", "))
		}

	}

	return sb.String() + bz.String()
}

const KEY = 811589153

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	i := 0
	list := make(List, 0)
	for fileScanner.Scan() {
		line = fileScanner.Text()
		n := Number{InitialPosition: i}
		fmt.Sscanf(line, "%d", &n.Val)
		list = append(list, n)
		i++
	}
	list2 := make(List, len(list))
	copy(list2, list)

	Mix(list)
	fmt.Println("Part 1:", Answer(list))

	for i := range list2 {
		list2[i].Val *= KEY
	}

	for i := 0; i < 10; i++ {
		Mix(list2)
		fmt.Println("Mix #", i, "List: ", list2)
	}

	fmt.Println("Part 2:", Answer(list2))

}

func Mix(list List) {
	// fmt.Println("Start", list)
	for i := 0; i < len(list); i++ {
		p := Find(i, list)
		Move(p, list)
		// fmt.Println(list)
	}
}

func Find(ip int, list List) int {
	for i, n := range list {
		if n.InitialPosition == ip {
			return i
		}
	}
	panic("uh oh")
}

func Move(p int, list List) {
	L := len(list)
	C := L - 1
	n := list[p]
	m := n.Val % C

	fp := (p + m) % L
	if m < 0 {
		fp += C
		fp = fp % L
	}
	limit := fp
	if fp < p {
		limit += L
	}

	// fmt.Printf("Moving %d from %d to pos %d between %d and %d\n", m, p, fp, list[(fp)%L].Val, list[(fp+1)%L].Val)

	for i := p; i < limit; i++ {
		list[i%L] = list[(i+1)%L]
	}

	list[fp] = n
}

func Answer(list List) int {
	for i, n := range list {
		if n.Val == 0 {
			L := len(list)
			a, b, c := i+1000, i+2000, i+3000
			fmt.Println(list[a%L].Val, list[b%L].Val, list[c%L].Val)
			return list[a%L].Val + list[b%L].Val + list[c%L].Val
		}
	}
	panic("uh oh")
}
