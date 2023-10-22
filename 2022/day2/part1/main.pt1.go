package part1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var MOVE_LIST = map[string]int{
	"A": 1, // Rock
	"B": 2, // Paper
	"C": 3, // Scissors
	"X": 1, // Rock
	"Y": 2, // Paper
	"Z": 3, // Scissors
}

type Move struct {
	name  string
	value int
}

type Round struct {
	Move1 Move
	Move2 Move
}

func NewMove(s string) Move {
	return Move{name: s, value: MOVE_LIST[s]}
}

func (r *Round) Score() (int, int) {
	var score [2]int

	switch true {

	case r.Move1.value == r.Move2.value:
		score[0] = 3
		score[1] = 3

	case (r.Move1.value - 1) == r.Move2.value%3:
		score[0] = 6
		score[1] = 0

	default:
		score[0] = 0
		score[1] = 6

	}

	score[0] += r.Move1.value
	score[1] += r.Move2.value
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
			Move1: NewMove(moves[0]),
			Move2: NewMove(moves[1]),
		}

		s1, s2 := r.Score()
		totalScore1 += s1
		totalScore2 += s2

		fmt.Println(moves, s1, s2)
	}
	fmt.Println(totalScore1, totalScore2)
}
