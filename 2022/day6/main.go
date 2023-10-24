package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileReader := bufio.NewReader(inputFile)

	const size int = 14

	buf := make([]rune, size, size)
	for i := 1; true; i++ {
		c, _, _ := fileReader.ReadRune()
		buf = append([]rune{c}, buf[:size-1]...)
		if i >= size && checkUnique(buf) {
			fmt.Println("end", string(c), buf, i)
			break
		}
		fmt.Println(string(c), buf, i)
	}
}

func checkUnique(buf []rune) bool {

	for i, v := range buf[:len(buf)-1] {
		for _, w := range buf[i+1:] {
			if v == w {
				return false
			}
		}
	}

	return true
}
