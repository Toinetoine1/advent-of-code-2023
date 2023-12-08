package main

import (
	_ "embed"
	"errors"
	"fmt"
	"slices"
	"strings"
)

type Point struct {
	value  byte
	number *Number
}

type Number struct {
	raw, column int
	length      int
	value       int
}

type Star struct {
	raw, column int
}

//go:embed input.txt
var schem string
var border [][]int

func main() {
	board := make([][]Point, len(strings.Split(schem, "\n")))
	fillBoard(&board)

	numbers := make([]Number, 0)
	border = [][]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
		{-1, -1},
		{1, 1},
		{1, -1},
		{-1, 1},
	}

	stars := make([]Star, 0)

	fillNumbers(&board, &numbers)
	filterNumbers(&board, &numbers)

	findStar(&board, &stars)

	step2 := 0

	for _, star := range stars {
		foundNumbers := make([]Number, 0)

		for _, delta := range border {
			newRaw := star.raw + delta[0]
			newColumn := star.column + delta[1]

			point := board[newRaw][newColumn]
			if point.number != nil && !slices.Contains(foundNumbers, *point.number) {
				foundNumbers = append(foundNumbers, *point.number)
			}
		}

		if len(foundNumbers) == 2 {
			step2 += foundNumbers[0].value * foundNumbers[1].value
		}
	}

	fmt.Println("Step 2:", step2)
}

func findStar(board *[][]Point, stars *[]Star) {
	for raw, e := range *board {

		for column, _ := range e {
			if (*board)[raw][column].value == '*' {
				*stars = append(*stars, Star{raw: raw, column: column})
			}
		}
	}
}

func filterNumbers(board *[][]Point, numbers *[]Number) {
	res := make([]Number, 0)

	for _, number := range *numbers {

	nb:
		for i := 0; i < number.length; i++ {
			for _, delta := range border {
				newRaw := number.raw + delta[0]
				newColumn := number.column + i + delta[1]

				if newRaw >= 0 && newColumn >= 0 && newRaw < len(*board) && newColumn < len((*board)[0]) {
					char := (*board)[newRaw][newColumn].value

					if !(char >= '0' && char <= '9') && char != '.' {
						res = append(res, number)
						break nb
					}
				}
			}
		}

	}

	sum := 0
	for _, number := range res {
		sum += number.value
	}

	fmt.Println("Step 1:", sum)
}

func fillNumbers(board *[][]Point, numbers *[]Number) {
	for raw := 0; raw < len(*board); raw++ {
		for column := 0; column < len((*board)[0]); column++ {
			number, err := parseNumber(board, raw, column)

			if err != nil {
				continue
			}

			*numbers = append(*numbers, number)
			column += number.length
		}
	}
}

func printBoard(board *[][]byte) {
	for _, e := range *board {

		for _, e2 := range e {
			fmt.Printf("%c", e2)
		}

		fmt.Println()
	}
}

func fillBoard(board *[][]Point) {
	array := strings.Split(schem, "\n")

	for i := range *board {
		line := array[i]

		raw := make([]Point, len(line))

		for j := range raw {
			raw[j] = Point{value: line[j]}
		}

		(*board)[i] = raw
	}
}

func parseNumber(board *[][]Point, raw int, index int) (Number, error) {
	number := Number{
		raw: raw,
	}

	line := (*board)[raw]
	columnSet := false

	for i := index; i < len(line); i++ {
		char := line[i].value

		if char >= '0' && char <= '9' {
			if !columnSet {
				columnSet = true
				number.column = i
			}

			line[i].number = &number
			number.value = number.value*10 + (int(char) - '0')
			number.length += 1
		} else if columnSet {
			return number, nil
		} else {
			return number, errors.New("no number here")
		}
	}

	if columnSet {
		return number, nil
	}

	return number, errors.New("no number here")
}
