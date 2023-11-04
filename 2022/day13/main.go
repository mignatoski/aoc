package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Pair struct {
	Left, Right         *Packet
	LeftText, RightText string
	Index               int
}

type Packet struct {
	List   []*Packet
	Value  int
	Parent *Packet
}

type PacketType int

const (
	NUMBER PacketType = iota
	LIST
)

func (pair *Pair) String() string {
	return fmt.Sprintf("-- PAIR %v --\n%v\n%v\n", pair.Index, pair.LeftText, pair.RightText)
}

func (packet *Packet) String() string {
	return fmt.Sprintf("-- PACKET --\nList: %v\nValue: %v\n", len(packet.List), packet.Value)
}

func (packet *Packet) Type() PacketType {
	if packet.Value < 0 {
		return LIST
	} else {
		return NUMBER
	}
}

type CompareResult int

const (
	CORRECT CompareResult = iota
	WRONG
	CONTINUE
)

func Compare(left, right *Packet) CompareResult {
	/*
	   If both values are integers, the lower integer should come first.
	   If the left integer is lower than the right integer, the inputs are in the right order.
	   If the left integer is higher than the right integer, the inputs are not in the right order.
	   Otherwise, the inputs are the same integer; continue checking the next part of the input.

	   If both values are lists, compare the first value of each list, then the second value, and so on.
	   If the left list runs out of items first, the inputs are in the right order.
	   If the right list runs out of items first, the inputs are not in the right order.
	   If the lists are the same length and no comparison makes a decision about the order,
	   continue checking the next part of the input.

	   If exactly one value is an integer, convert the integer to a list which contains that integer as its only value,
	   then retry the comparison. For example, if comparing [0,0,0] and 2, convert the right value to [2]
	   (a list containing 2); the result is then found by instead comparing [0,0,0] and [2].
	*/
	switch {

	case left.Type() == NUMBER && right.Type() == NUMBER:
		if left.Value < right.Value {
			return CORRECT
		} else if left.Value > right.Value {
			return WRONG
		} else {
			return CONTINUE
		}

	case left.Type() == LIST && right.Type() == LIST:
		lenL, lenR := len(left.List), len(right.List)
		for i := 0; i < min(lenL, lenR); i++ {

			switch Compare(left.List[i], right.List[i]) {

			case CORRECT:
				return CORRECT
			case WRONG:
				return WRONG
			case CONTINUE:
				continue
			}
		}

		switch {
		case lenL < lenR:
			return CORRECT
		case lenL > lenR:
			return WRONG
		default:
			return CONTINUE
		}

	case left.Type() == NUMBER && right.Type() == LIST:
		if len(right.List) == 0 {
			return WRONG
		}
		switch Compare(left, right.List[0]) {
		case CORRECT:
			return CORRECT
		case WRONG:
			return WRONG
		default:
			if len(right.List) > 1 {
				return CORRECT
			} else {
				return CONTINUE
			}
		}

	case left.Type() == LIST && right.Type() == NUMBER:
		if len(left.List) == 0 {
			return CORRECT
		}
		switch Compare(left.List[0], right) {
		case CORRECT:
			return CORRECT
		case WRONG:
			return WRONG
		default:
			if len(left.List) > 1 {
				return WRONG
			} else {
				return CONTINUE
			}
		}
	}

	panic("Identical pair, shouldn't happen")

}

func NewPacket() Packet {
	return Packet{List: make([]*Packet, 0), Value: -1}
}

func BuildPacket(text string) *Packet {
	packet := NewPacket()

	cur := &packet
	chars := []rune(text)

	for i := 0; i < len(chars); i++ {
		switch chars[i] {

		case '[':
			p := NewPacket()
			p.Parent = cur
			cur.List = append(cur.List, &p)
			cur = &p

		case ']':
			cur = cur.Parent

		case ',':
			// do nothing

		default:
			// deal with two digit numbers
			numText := string(chars[i])
			if i < len(chars)-1 && chars[i+1] >= '0' && chars[i+1] <= '9' {
				numText += string(chars[i+1])
				i++
			}
			number, _ := strconv.ParseInt(string(numText), 0, 0)
			p := NewPacket()
			p.Value = int(number)
			p.Parent = cur
			cur.List = append(cur.List, &p)

		}

	}

	// fmt.Println(packet)

	return &packet
}

func main() {
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	lineCount := 0
	indexSum := 0

	pairs := make([]*Pair, 0)
	var pair *Pair

	for fileScanner.Scan() {
		line = fileScanner.Text()

		switch lineCount % 3 {
		case 0:
			pair = &Pair{}
			pair.LeftText = line
			pair.Index = (lineCount / 3) + 1
			pair.Left = BuildPacket(line[1 : len(line)-1])
			pairs = append(pairs, pair)

		case 1:
			pair.RightText = line
			pair.Right = BuildPacket(line[1 : len(line)-1])
		case 2:
			// do nothing, blank line
			fmt.Println(pair)
			result := Compare(pair.Left, pair.Right)
			fmt.Println(result)

			if result == CORRECT {
				indexSum += pair.Index
			}
		}

		lineCount++
	}

	fmt.Println("Part 1 Sum of Indexes: ", indexSum)

	p := Pair{}
	p.Left = BuildPacket("[[[]]]")
	p.Right = BuildPacket("[[[[[]]]]]")

	fmt.Println(Compare(p.Left, p.Right))

}
