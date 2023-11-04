package main

import (
	"fmt"
	"os"

	"github.com/Haugen/aoc-go/internal/utils"
)

func main() {
	path, _ := os.Getwd()
	data := utils.FilenameToArray(path + "/input.txt")

	part1(data)
	part2(data)
}

func part1(data []string) {
	var count int

	for i, line := range data[:len(data)-1] {
		if utils.StrToInt(data[i+1]) > utils.StrToInt(line) {
			count++
		}
	}

	fmt.Println(count)
}

func part2(data []string) {
	count := 0

	for i := 0; i < len(data[:len(data)-3]); i++ {
		curr := data[i : i+3]
		next := data[i+1 : i+4]

		currCount := utils.SumArray(utils.StrArrayToIntArray(curr))
		nextCount := utils.SumArray(utils.StrArrayToIntArray(next))

		if nextCount > currCount {
			count++
		}
	}

	fmt.Println(count)
}
