package main

import (
	"bufio"
	"fmt"
	"os"
)

type Number struct {
	Val, InitialPosition int
}

const KEY = 811589153

func main() {
	inputFile, _ := os.Open("sample.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	i := 0
	list := make([]Number, 0)
	for fileScanner.Scan() {
		line = fileScanner.Text()
		n := Number{InitialPosition: i}
		fmt.Sscanf(line, "%d", &n.Val)
		list = append(list, n)
		i++
	}
	list2 := make([]Number, len(list))
	copy(list2, list)

	Mix(list)
	fmt.Println("Part 1:", Answer(list))

	for i := range list2 {
		list2[i].Val *= KEY
	}

	for i := 0; i < 10; i++ {
		Mix(list2)
		fmt.Println("Mix #", i)
	}

	fmt.Println("Part 2:", Answer(list2))

}

func Mix(list []Number) {
	for i := 0; i < len(list); i++ {
		// fmt.Println(list)
		p := Find(i, list)
		Move(p, list)
	}
}

func Find(ip int, list []Number) int {
	for i, n := range list {
		if n.InitialPosition == ip {
			return i
		}
	}
	panic("uh oh")
}

func Move(p int, list []Number) {
	L := len(list)
	n := list[p]

	m := n.Val % (L)
	adj := n.Val / L
	fp := ((p + m) % L) + adj
	for fp >= L {
		m = (m + adj) % L
		adj = (m + adj) / L
		fp = ((p + m) % L) + adj
	}
	for fp < 0 {
		fp += L - 1
	}
	limit := fp
	if fp < p {
		limit += L
	}
	// fmt.Println(n, p, fp, L, limit)

	for i := p; i < limit; i++ {
		list[i%L] = list[(i+1)%L]
	}

	list[fp] = n
}

func Answer(list []Number) int {
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
