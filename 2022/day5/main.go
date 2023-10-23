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

	var line, section string

	crateData := make([]string, 0)
	actionData := make([]string, 0)
	section = "crates"

	for fileScanner.Scan() {
		line = fileScanner.Text()
		if strings.Contains(line, " 1   2   3   4   5") {
			section = "actions"
			continue
		}

		if line == "" {
			continue
		}

		switch section {

		case "crates":
			crateData = append(crateData, line)

		case "actions":
			actionData = append(actionData, line)

		}
	}

	stacks := make([][]string, 9, 9)

	for _, l := range crateData {
		for i := 0; i < 9; i++ {
			v := string(l[1+(4*i)])
			if v != " " {
				stacks[i] = append(stacks[i], v)
			}
		}
	}

	fmt.Println(stacks)
	// fmt.Println(move(&stacks[0], &stacks[1]))

	for _, l := range actionData {
		action := strings.Split(l, " ")
		count, _ := strconv.ParseInt(action[1], 10, 0)
		from, _ := strconv.ParseInt(action[3], 10, 0)
		to, _ := strconv.ParseInt(action[5], 10, 0)

		fmt.Println(count, from, to)
		move2(&stacks[from-1], &stacks[to-1], int(count))
		fmt.Println(stacks)
	}

	fmt.Println(stacks)

	for _, s := range stacks {
		fmt.Print(s[0])
	}

}

func move2(f *[]string, t *[]string, count int) {
	result := make([]string, count)
	// fmt.Println(result, *f, *t)
	copy(result, (*f)[:count])
	*f = (*f)[count:]
	*t = append(result, (*t)...)
	// fmt.Println(result, *f, *t)
}
func move(f *[]string, t *[]string) string {
	var result string
	result, *f = (*f)[0], (*f)[1:]
	*t = append([]string{result}, (*t)...)
	return result
}
