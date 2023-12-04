package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/Haugen/aoc-go/internal/utils"
)

var data []string

func main() {
	path, _ := os.Getwd()
	data = utils.FilenameToArray(path + "/input.txt")

	part1()
	part2()
}

func getWinningCount(winners []string, numbers []string) int {
	cardCount := 0
	for _, winningNr := range winners {
		if winningNr == "" {
			continue
		}
		if slices.Contains(numbers, winningNr) {
			cardCount++
		}
	}
	return cardCount
}

func part1() {
	total := 0

	for _, line := range data {
		splitLine := strings.Split(line, ":")
		_, cards := strings.Trim(splitLine[0][5:], " "), strings.Split(splitLine[1], "|")
		winning, mine := strings.Split(cards[0], " "), strings.Split(cards[1], " ")
		cardCount := getWinningCount(winning, mine)
		multiple := 0
		for i := 0; i < cardCount; i++ {
			if multiple == 0 {
				multiple = 1
			} else {
				multiple = multiple * 2
			}
		}
		total += multiple
	}

	fmt.Println(total)
}

// Terrible performance, but works
func part2() {
	cardCounter := make([]int, 211)
	for i := range cardCounter {
		cardCounter[i] = 1
	}
	var cards [][]string

	for _, line := range data {
		splitLine := strings.Split(line, ":")
		cards = append(cards, strings.Split(splitLine[1], "|"))
	}

	for i := 0; i < len(cards); i++ {
		for j := 0; j < cardCounter[i]; j++ {
			cardCount := 0

			winning, mine := strings.Split(cards[i][0], " "), strings.Split(cards[i][1], " ")
			cardCount = getWinningCount(winning, mine)

			for k := 1; k <= cardCount; k++ {
				cardCounter[i+k]++
			}
		}
	}

	total := 0
	for _, num := range cardCounter {
		total += num
	}
	fmt.Println(total)
}
