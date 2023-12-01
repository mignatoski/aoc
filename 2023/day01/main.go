package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numbers map[string]rune

func init() {

	numbers = map[string]rune{
		// "zero":  '0',
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	// fmt.Println(ReplaceDigitsLast("sevenninethreefive4bgknpbnine"))
}

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	total, total2 := 0, 0
	for fileScanner.Scan() {
		line = fileScanner.Text()

		// part 1

		runes := []rune(line)
		n := len(line)
		fd, sd := '0', '0'

		for i := 0; i < n; i++ {
			r := runes[i]
			if r >= '0' && r <= '9' {
				fd = r
				break
			}
		}

		for i := n - 1; i >= 0; i-- {
			r := runes[i]
			if r >= '0' && r <= '9' {
				sd = r
				break
			}
		}

		str := string(fd) + string(sd)

		num, err := strconv.ParseInt(str, 0, 0)

		if err != nil || len(str) != 2 {
			panic(str)
		}

		total += int(num)

		// part 2
		orig := line

		line = ReplaceDigits(orig)
		runes = []rune(line)
		n = len(runes)
		fd = 'a'

		for i := 0; i < n; i++ {
			r := runes[i]
			if r >= '0' && r <= '9' {
				fd = r
				break
			}
		}

		line = ReplaceDigitsLast(orig)
		runes = []rune(line)
		n = len(runes)
		sd = 'a'

		for i := n - 1; i >= 0; i-- {
			r := runes[i]
			if r >= '0' && r <= '9' {
				sd = r
				break
			}
		}

		str = string(fd) + string(sd)

		num, err = strconv.ParseInt(str, 0, 0)

		if err != nil || len(str) != 2 {
			panic(str)
		}
		fmt.Println(int(num))

		total2 += int(num)
	}
	fmt.Println("Part 1: ", total)
	fmt.Println("Part 2: ", total2)

	fmt.Println(ReplaceDigitsLast("sevenninethreefive4bgknpbnine"))
}

func ReplaceDigits(s string) string {
	fmt.Println(s)
	first := 999999
	digit := ""

	for k := range numbers {
		m := strings.Index(s, k)
		if m >= 0 && m < first {
			first = m
			digit = k
		}
	}

	if digit != "" {
		s = strings.Replace(s, digit, string(numbers[digit]), 1)
	}
	return s
}

func ReplaceDigitsLast(s string) string {
	last := -1
	digit := ""

	for k := range numbers {
		m := strings.LastIndex(s, k)
		if m >= 0 && m > last {
			last = m
			digit = k
		}
	}

	if digit != "" {
		s = strings.ReplaceAll(s, digit, string(numbers[digit]))
	}
	return s
}
