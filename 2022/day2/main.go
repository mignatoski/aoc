package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var MOVE_LIST = map[string]int{
	"A": 0, // Rock
	"B": 1, // Paper
	"C": 2, // Scissors
}

var RESULT_LIST = map[string]int{
	"X": 2, // Lose
	"Y": 0, // Draw
	"Z": 1, // Win
}

type Move struct {
	name  string
	value int
}

type Round struct {
	Move1  Move
	Result Move
}

func NewMove(s string) Move {
	return Move{name: s, value: MOVE_LIST[s]}
}

func NewResult(s string) Move {
	return Move{name: s, value: RESULT_LIST[s]}
}

func (r *Round) Score() (int, int) {
	var score [2]int

	switch true {

	case r.Result.value == 0: // Draw
		score[0] = 3 + r.Move1.value + 1
		score[1] = 3 + r.Move1.value + 1

	case r.Result.value == 2: // Lose
		score[0] = 6 + r.Move1.value + 1
		score[1] = 0 + ((r.Move1.value + r.Result.value) % 3) + 1

	default: // Win
		score[0] = 0 + r.Move1.value + 1
		score[1] = 6 + ((r.Move1.value + r.Result.value) % 3) + 1

	}

	return score[0], score[1]
}

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	var moves []string
	var totalScore1, totalScore2 int
	for fileScanner.Scan() {
		line = fileScanner.Text()
		moves = strings.Split(line, " ")
		r := Round{
			Move1:  NewMove(moves[0]),
			Result: NewResult(moves[1]),
		}

		s1, s2 := r.Score()
		totalScore1 += s1
		totalScore2 += s2

		fmt.Println(moves, s1, s2)
	}
	fmt.Println(totalScore1, totalScore2)
}
