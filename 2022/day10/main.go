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

	var line string

	cycle, x := 1, 1
	stack := make([]int, 0)
	isEmpty := true
	totalSignalStrength := 0
	crt := ""

	for {
		// Start Cycle
		isEmpty = len(stack) == 0

		if isEmpty {

			if !fileScanner.Scan() {
				break
			}
			line = fileScanner.Text()
			inst, v, _ := strings.Cut(line, " ")
			fmt.Println(cycle, x, inst, v)

			switch inst {
			case "noop":
			case "addx":
				intV, _ := strconv.ParseInt(v, 0, 0)
				stack = append(stack, int(intV))
			}

		}

		// During Cycle
		if (cycle+20)%40 == 0 {
			totalSignalStrength += x * cycle
		}

		horizontalPos := (cycle - 1) % 40

		if horizontalPos >= x-1 && horizontalPos <= x+1 {
			crt += "#"
		} else {
			crt += "."
		}
		if horizontalPos == 39 {
			crt += "\n"
		}

		// End Of Cycle
		if !isEmpty {
			fmt.Println(cycle, x, stack)
			x += stack[0]
			stack = stack[1:]
		}
		cycle++
	}

	fmt.Println(totalSignalStrength)
	fmt.Println(crt)
}
