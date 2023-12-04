package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var input string = "Game 1: 1 red, 3 blue, 11 green; 1 blue, 5 red; 3 blue, 5 green, 13 red; 6 red, 1 blue, 4 green; 16 red, 12 green\nGame 2: 3 red, 13 blue, 5 green; 14 green, 14 blue; 9 blue, 10 green, 3 red; 2 green, 5 blue; 11 green, 3 blue, 3 red; 16 blue, 2 red, 9 green\nGame 3: 17 blue, 5 red; 3 red, 11 green, 17 blue; 1 red, 6 blue, 9 green; 3 blue, 11 green, 1 red; 3 green, 10 red, 11 blue; 12 red, 3 green, 15 blue\nGame 4: 14 green, 14 red, 1 blue; 15 red, 13 green, 1 blue; 6 green, 15 red; 7 green\nGame 5: 3 green, 1 blue, 3 red; 6 red, 2 green, 2 blue; 12 red, 3 green, 1 blue; 2 green, 9 red; 1 blue; 2 blue, 10 red\nGame 6: 5 blue, 5 green; 4 blue, 1 red, 10 green; 16 green, 1 red, 6 blue; 1 red, 1 blue, 13 green; 1 red, 5 blue, 7 green; 14 green, 17 blue\nGame 7: 1 green, 8 blue, 4 red; 1 green, 4 blue, 4 red; 6 blue, 4 red, 4 green; 1 red, 8 green\nGame 8: 2 red, 5 blue, 1 green; 1 blue, 4 red, 8 green; 6 blue, 12 green, 6 red; 3 blue, 5 red; 8 red, 2 blue, 13 green; 5 green, 4 red, 3 blue\nGame 9: 11 red; 1 green, 2 red, 2 blue; 1 blue, 2 green, 9 red; 4 red, 2 green, 2 blue; 1 blue, 2 green; 1 blue, 9 red, 2 green\nGame 10: 9 red, 4 green; 1 blue, 3 red, 7 green; 3 green, 1 red, 1 blue; 7 green, 4 red, 1 blue; 1 blue, 5 green, 10 red; 1 red, 5 green\nGame 11: 2 blue, 4 red, 3 green; 1 blue, 7 red; 4 green, 7 red, 1 blue; 3 blue, 6 green, 4 red; 3 red, 1 green, 3 blue\nGame 12: 1 green, 6 red, 5 blue; 3 green, 2 red, 4 blue; 3 green, 1 red, 3 blue\nGame 13: 6 green, 1 red, 9 blue; 11 red, 4 blue, 12 green; 6 green, 9 red, 19 blue; 2 green, 6 blue; 10 green, 1 red, 16 blue; 4 green, 14 blue\nGame 14: 7 blue, 2 red; 1 green, 2 red, 19 blue; 12 blue, 6 green, 11 red\nGame 15: 4 red, 4 green, 7 blue; 15 blue, 1 green, 8 red; 2 red, 10 green, 11 blue; 5 red, 4 blue, 6 green; 9 red, 8 blue, 3 green; 9 blue, 9 red\nGame 16: 7 red, 2 blue, 19 green; 6 blue, 9 green; 8 green, 6 red, 19 blue; 11 green, 7 red, 1 blue; 9 blue, 3 red, 17 green\nGame 17: 3 blue, 4 green, 5 red; 2 red, 4 green, 11 blue; 6 blue, 13 green; 3 blue, 12 green, 7 red\nGame 18: 9 red, 6 blue, 7 green; 3 green, 3 blue, 5 red; 18 red, 6 blue, 4 green; 3 green, 10 red, 8 blue\nGame 19: 3 red, 6 green; 1 red, 5 green, 4 blue; 3 red, 14 blue\nGame 20: 2 green, 2 blue, 4 red; 14 red, 6 blue, 5 green; 1 blue, 5 red, 3 green; 10 red, 6 green, 6 blue\nGame 21: 10 blue, 12 green, 3 red; 1 green, 14 red; 5 blue, 7 green; 12 blue, 1 red, 13 green; 7 red, 4 green\nGame 22: 2 red, 1 blue; 1 red, 2 blue; 1 red, 1 green, 3 blue; 3 blue; 1 red; 1 green, 2 red\nGame 23: 4 blue, 4 green, 1 red; 3 blue, 1 red, 6 green; 1 red, 1 blue\nGame 24: 5 blue, 15 green, 13 red; 20 green, 13 blue, 6 red; 5 blue, 11 red, 16 green; 6 red, 5 blue, 13 green; 12 blue, 13 green, 3 red\nGame 25: 10 blue, 17 red; 12 red, 16 blue, 3 green; 4 green, 12 blue, 10 red; 8 blue, 3 green, 10 red; 5 green, 2 red, 12 blue\nGame 26: 11 red, 9 blue; 3 blue, 3 red, 3 green; 10 blue, 3 green, 4 red; 1 green, 4 blue, 9 red; 5 green, 1 red, 7 blue; 1 red, 3 blue, 3 green\nGame 27: 1 green, 12 red, 4 blue; 5 red, 2 green, 1 blue; 3 green, 6 blue, 10 red; 1 green, 4 red, 3 blue\nGame 28: 6 blue; 2 green; 2 green, 8 blue, 1 red; 2 green, 2 blue; 6 blue, 8 green; 9 green, 5 blue\nGame 29: 1 green, 9 blue, 9 red; 13 green, 4 red, 9 blue; 3 green, 8 blue, 15 red; 15 green, 18 blue, 3 red; 16 green, 10 red; 16 green, 12 blue, 16 red\nGame 30: 14 blue, 4 green, 1 red; 7 red, 14 blue; 2 blue, 4 red, 1 green\nGame 31: 2 red, 14 green, 3 blue; 3 blue, 3 green, 4 red; 8 blue, 4 red, 1 green; 8 green, 3 blue; 10 blue, 1 red, 11 green; 13 green, 2 red, 3 blue\nGame 32: 8 blue, 16 red; 2 green, 8 blue, 16 red; 16 blue, 4 green, 17 red; 2 red, 5 green, 4 blue\nGame 33: 2 red, 2 green, 1 blue; 5 red, 1 blue; 8 green, 14 red\nGame 34: 4 red, 4 green; 9 green; 1 blue, 16 green; 1 blue, 5 red, 9 green; 2 red, 15 green, 1 blue\nGame 35: 1 green, 5 red; 1 green, 15 red, 13 blue; 2 red, 13 blue, 17 green; 9 blue, 3 red, 11 green; 7 green, 8 blue, 14 red\nGame 36: 19 green; 3 green, 1 blue, 1 red; 1 green, 8 blue; 13 green, 5 red, 5 blue\nGame 37: 12 red, 7 green, 3 blue; 12 blue, 10 red, 9 green; 17 green, 8 red, 13 blue; 9 blue, 9 green, 8 red; 4 red, 13 green, 13 blue; 15 green, 12 red, 14 blue\nGame 38: 5 blue, 1 green, 20 red; 1 green, 13 red, 18 blue; 17 blue, 9 red, 10 green; 4 blue, 4 red, 12 green; 12 blue, 12 red, 6 green; 12 green, 13 red, 2 blue\nGame 39: 7 blue, 6 red, 2 green; 6 blue, 1 red; 7 blue, 1 red\nGame 40: 1 blue, 3 red; 15 blue, 1 green; 1 green, 16 red, 2 blue\nGame 41: 2 blue, 4 green; 8 green, 3 red; 2 blue, 9 red, 4 green; 4 red, 3 blue, 10 green; 5 green, 3 blue, 2 red\nGame 42: 7 green, 2 blue, 1 red; 8 green, 4 red; 5 blue, 1 red, 3 green\nGame 43: 3 red, 1 blue; 1 blue, 2 green, 2 red; 1 red, 2 blue; 3 blue\nGame 44: 3 green, 14 blue, 1 red; 16 blue, 5 red, 11 green; 12 green, 1 blue; 13 blue, 1 red; 5 blue, 2 red, 6 green; 3 blue, 5 red, 11 green\nGame 45: 7 blue, 1 red; 1 red, 3 blue; 3 green, 14 blue, 1 red; 4 blue, 3 green, 1 red; 15 blue, 1 red, 3 green\nGame 46: 15 red, 4 blue; 15 red, 11 blue, 3 green; 14 red, 2 green, 2 blue; 14 red, 8 blue, 3 green; 4 red, 1 blue\nGame 47: 4 green, 2 blue, 3 red; 8 red, 2 green, 18 blue; 1 green, 17 blue, 1 red\nGame 48: 2 green, 4 red, 2 blue; 15 blue, 16 red, 5 green; 14 blue, 2 green, 10 red; 3 green, 13 red, 6 blue; 8 green, 4 red, 12 blue; 15 red, 3 green, 9 blue\nGame 49: 1 green, 6 red, 7 blue; 1 blue, 9 green, 9 red; 4 green, 8 red; 9 blue, 1 red, 14 green; 2 blue, 9 red\nGame 50: 3 red, 10 blue, 14 green; 2 red, 9 blue, 7 green; 4 blue, 12 green; 1 red, 4 green, 5 blue\nGame 51: 2 green, 6 blue; 1 green, 10 blue, 1 red; 3 blue, 2 green\nGame 52: 1 green, 4 red, 1 blue; 3 red, 5 green, 4 blue; 1 blue, 3 red, 5 green; 1 red, 1 green, 1 blue; 12 green, 2 red, 4 blue; 10 blue, 7 green, 1 red\nGame 53: 12 red, 1 blue; 8 red, 11 blue, 11 green; 8 red, 6 blue, 13 green; 11 blue, 11 red, 16 green; 6 red, 9 green, 4 blue\nGame 54: 2 red, 8 blue, 15 green; 4 green, 3 blue, 6 red; 12 green, 13 blue, 4 red\nGame 55: 1 green, 16 blue, 4 red; 3 red, 1 blue, 1 green; 12 red, 16 blue; 3 red\nGame 56: 4 green; 1 red, 4 green; 2 red, 3 blue, 7 green; 2 red, 3 blue, 15 green\nGame 57: 17 green; 1 green, 9 blue; 1 red, 1 green, 9 blue\nGame 58: 3 green, 8 red, 7 blue; 4 green, 9 blue, 2 red; 1 red, 2 green, 11 blue; 8 blue, 4 green\nGame 59: 6 green, 1 red; 4 blue, 6 green; 4 green, 5 blue\nGame 60: 3 green, 5 blue, 1 red; 7 green, 5 blue, 16 red; 14 red, 1 green, 1 blue; 7 green, 2 blue; 13 red, 5 green, 5 blue\nGame 61: 1 green, 2 blue, 2 red; 2 green; 6 red, 1 blue, 1 green\nGame 62: 5 red, 8 blue, 1 green; 1 red, 1 blue; 2 green, 8 blue\nGame 63: 2 red, 2 blue, 2 green; 9 blue, 7 green; 1 green, 4 blue; 18 green, 3 blue\nGame 64: 13 green, 1 blue, 6 red; 13 green, 15 red, 8 blue; 5 green, 14 red, 4 blue; 2 green, 8 blue, 12 red; 1 blue, 5 red, 13 green; 7 blue, 8 green, 2 red\nGame 65: 7 blue, 12 red, 6 green; 11 red, 8 green, 8 blue; 9 red, 7 green, 7 blue; 14 red, 2 blue, 17 green\nGame 66: 2 green, 5 red; 7 red, 14 blue; 19 blue, 2 green; 7 blue, 4 green, 6 red\nGame 67: 4 green, 17 red, 7 blue; 4 blue, 6 green; 7 green, 7 red, 12 blue; 2 red, 14 blue\nGame 68: 1 red, 11 green, 4 blue; 17 blue, 1 red, 10 green; 3 red, 7 green, 1 blue; 7 green, 3 red, 6 blue; 2 red, 3 green; 2 green, 2 red, 4 blue\nGame 69: 5 blue, 4 red; 3 red, 11 green, 1 blue; 6 green, 2 blue; 10 green, 4 red, 5 blue; 2 red, 11 green\nGame 70: 16 red, 7 blue, 1 green; 14 red, 1 blue, 4 green; 4 red, 4 green; 7 blue, 5 red, 2 green\nGame 71: 14 red, 2 blue, 13 green; 7 green, 5 red, 2 blue; 3 blue, 9 green, 11 red; 10 red, 4 blue, 1 green\nGame 72: 1 green, 2 red, 6 blue; 4 green, 4 red, 9 blue; 6 green, 8 blue, 1 red; 5 red, 4 green, 9 blue; 15 blue, 2 green, 7 red; 10 blue, 2 green, 10 red\nGame 73: 7 green, 6 red, 7 blue; 6 blue, 5 red, 8 green; 5 blue, 5 red\nGame 74: 11 red, 1 blue; 2 green, 4 blue, 1 red; 1 green, 2 blue, 11 red; 9 red, 5 blue; 15 red, 10 blue; 9 red, 3 blue\nGame 75: 1 blue, 6 red, 9 green; 5 red, 1 blue, 8 green; 2 green, 2 red, 1 blue; 7 red, 1 green; 3 green, 6 red, 2 blue; 1 green, 1 red\nGame 76: 16 red, 3 blue, 9 green; 4 blue, 4 green; 5 blue, 1 green, 10 red; 6 blue, 13 red; 1 blue, 2 green, 8 red\nGame 77: 4 red, 4 blue; 5 blue, 5 red; 6 red, 3 green\nGame 78: 11 green, 1 red; 1 blue, 18 green, 1 red; 6 green, 5 red, 2 blue; 6 red, 1 blue, 15 green; 5 green, 5 red\nGame 79: 2 red, 3 green, 13 blue; 7 blue, 5 green; 4 blue, 2 red, 6 green; 6 green, 15 blue\nGame 80: 9 green, 2 blue, 1 red; 8 green, 1 red; 1 blue, 7 green; 2 green, 1 blue; 3 green; 5 green, 1 red, 2 blue\nGame 81: 2 blue, 8 green, 1 red; 3 green, 1 blue; 6 blue, 1 green; 3 blue, 3 green, 1 red; 2 green, 8 blue; 1 red, 8 blue, 2 green\nGame 82: 5 blue, 4 red, 1 green; 9 red, 12 green, 8 blue; 9 red, 6 green, 15 blue; 8 blue, 10 red, 6 green\nGame 83: 2 green, 7 blue, 4 red; 2 blue, 11 red, 9 green; 7 red, 7 green, 6 blue; 12 blue, 4 red, 11 green; 11 green, 7 blue; 7 green, 5 red, 2 blue\nGame 84: 9 red, 1 blue, 7 green; 5 red, 5 green; 4 green, 4 blue; 4 green, 5 red\nGame 85: 5 green, 13 red, 11 blue; 5 blue, 19 green, 15 red; 17 red, 3 green, 8 blue; 13 green, 10 red; 3 green, 17 red, 11 blue\nGame 86: 1 green, 11 blue; 11 blue, 1 green, 8 red; 6 blue, 4 red; 4 blue, 17 red; 1 green, 15 red\nGame 87: 3 green, 8 red, 6 blue; 6 red, 13 green, 1 blue; 4 blue, 8 red, 8 green\nGame 88: 5 green, 5 blue; 3 green, 10 blue, 2 red; 6 blue, 7 red, 1 green; 5 green, 3 red, 11 blue; 8 red, 4 green, 6 blue\nGame 89: 5 green, 10 blue, 12 red; 1 green, 13 red, 8 blue; 4 red, 11 green, 12 blue\nGame 90: 4 green, 3 red, 11 blue; 1 green, 12 red, 12 blue; 9 blue, 5 red, 1 green; 2 green, 12 blue, 12 red\nGame 91: 5 red, 8 blue, 1 green; 5 green, 3 blue; 9 blue, 7 green, 5 red; 1 green, 3 blue, 6 red; 9 blue, 11 green, 4 red; 2 green, 4 red, 10 blue\nGame 92: 11 blue, 1 red, 6 green; 10 blue, 2 red; 4 red, 6 green, 19 blue\nGame 93: 1 green, 3 blue, 3 red; 3 red; 5 blue, 3 red; 1 green, 4 red\nGame 94: 9 red, 4 blue, 4 green; 1 blue, 6 red, 15 green; 10 red, 5 blue, 1 green; 2 blue, 4 green, 8 red\nGame 95: 13 blue, 4 green, 3 red; 15 green, 3 red, 2 blue; 16 green, 8 blue, 2 red\nGame 96: 15 blue, 7 green, 3 red; 5 red, 7 green, 17 blue; 6 red, 12 blue; 5 green, 10 blue, 4 red\nGame 97: 5 red, 2 green; 8 red; 1 blue, 7 green, 2 red; 7 red, 15 green\nGame 98: 6 green, 1 blue, 1 red; 3 green, 3 red; 1 blue, 13 green, 4 red\nGame 99: 16 red, 5 blue, 9 green; 2 green, 7 blue, 2 red; 10 blue, 3 green; 9 red, 8 blue, 13 green; 16 green, 13 red, 10 blue\nGame 100: 16 blue, 12 red, 3 green; 2 green, 7 blue; 5 blue, 4 green; 10 blue, 6 red, 6 green; 5 red, 12 blue, 2 green; 9 red, 12 blue, 11 green"

type Game struct {
	gameNumber, maxRed, maxBlue, maxGreen, minRed, minBlue, minGreen int
}

func main() {
	inputArray := strings.Split(input, "\n")

	var games = make([]Game, len(inputArray))
	parseGames(inputArray, &games)

	challenge1, challenge2 := 0, 0

	for _, game := range games {
		if game.maxRed <= 12 && game.maxBlue <= 14 && game.maxGreen <= 13 {
			challenge1 += game.gameNumber
		}

		challenge2 += game.maxGreen * game.maxBlue * game.maxRed
	}

	fmt.Printf("Challenge 1: %d\n", challenge1)
	fmt.Printf("Challenge 2: %d\n", challenge2)
}

func parseGames(inputArray []string, games *[]Game) {
	r, _ := regexp.Compile(" ([0-9]+) ([a-z]+)")

	for i, e := range inputArray {
		game := &(*games)[i]

		game.gameNumber = i + 1
		game.minBlue = math.MaxInt
		game.minGreen = math.MaxInt
		game.minRed = math.MaxInt
		e = e[7:]
		sets := strings.Split(e, ";")
		var red = &(game.maxRed)
		var blue = &(game.maxBlue)
		var green = &(game.maxGreen)

		for _, set := range sets {
			colors := strings.Split(set, ",")

			for _, color := range colors {
				match := r.FindStringSubmatch(color)

				if "green" == match[2] {
					tmp, _ := strconv.Atoi(match[1])

					if tmp > *green {
						*green = tmp
					}
					if game.minGreen > tmp {
						game.minGreen = tmp
					}
				} else if "red" == match[2] {
					tmp, _ := strconv.Atoi(match[1])

					if tmp > *red {
						*red = tmp
					}
					if game.minRed > tmp {
						game.minRed = tmp
					}
				} else if "blue" == match[2] {
					tmp, _ := strconv.Atoi(match[1])

					if tmp > *blue {
						*blue = tmp
					}
					if game.minBlue > tmp {
						game.minBlue = tmp
					}
				} else {
					log.Fatal("No match: ", match[1])
				}
			}

		}

	}
}
