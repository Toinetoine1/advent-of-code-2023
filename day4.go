package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var schematic string

type Card struct {
	winningNumbers, scratchedNumbers []int
}

func main() {
	inputArray := strings.Split(schematic, "\n")

	cards := make([]Card, 0)

	fillCards(&inputArray, &cards)

	step1(cards)
	step2(cards)
}

func step1(cards []Card) {
	step1 := 0

	for _, card := range cards {
		match := 0

		for _, number := range card.scratchedNumbers {
			if slices.Contains(card.winningNumbers, number) {
				match += 1
			}
		}

		if match != 0 {
			i := powInt(2, match-1)
			step1 += i
		}
	}

	fmt.Println("Step 1:", step1)
}

func step2(cards []Card) {
	step2 := 0
	remainingsCards := make([]int, 0)

	for i := 0; i < len(cards); i++ {
		step2 += 1
		card := cards[i]
		match := 0

		for _, number := range card.scratchedNumbers {
			if slices.Contains(card.winningNumbers, number) {
				match += 1
			}
		}
		for j := 0; j < match; j++ {
			elems := i + 1 + j
			remainingsCards = append(remainingsCards, elems)
			step2 += 1
		}
	}

	for len(remainingsCards) != 0 {
		i := remainingsCards[0]
		remainingsCard := cards[i]
		match := 0

		for _, number := range remainingsCard.scratchedNumbers {
			if slices.Contains(remainingsCard.winningNumbers, number) {
				match += 1
			}
		}

		remainingsCards = remainingsCards[1:]

		for j := 0; j < match; j++ {
			elems := i + 1 + j
			remainingsCards = append(remainingsCards, elems)
			step2 += 1
		}
	}

	fmt.Println("Step 2:", step2)
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func fillCards(inputArray *[]string, cards *[]Card) {
	for _, line := range *inputArray {
		line = strings.Split(line, ": ")[1]
		array := strings.Split(line, " | ")

		winningNumbers := make([]int, 0)
		scratchedNumbers := make([]int, 0)

		for loopIndex := 0; loopIndex < len(array[0]); loopIndex += 3 {
			winningNumber, _ := strconv.Atoi(strings.TrimSpace(string(array[0][loopIndex]) + string(array[0][loopIndex+1])))

			winningNumbers = append(winningNumbers, winningNumber)
		}

		for loopIndex := 0; loopIndex < len(array[1]); loopIndex += 3 {
			winningNumber, _ := strconv.Atoi(strings.TrimSpace(string(array[1][loopIndex]) + string(array[1][loopIndex+1])))

			scratchedNumbers = append(scratchedNumbers, winningNumber)
		}

		card := Card{winningNumbers: winningNumbers, scratchedNumbers: scratchedNumbers}

		*cards = append(*cards, card)
	}

}
