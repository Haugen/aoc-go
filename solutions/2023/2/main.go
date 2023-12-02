package main

import (
	"fmt"
	"os"
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

func part1() {
	sum := 0

outer:
	for _, line := range data {
		splitLine := strings.Split(line, ":")
		game, sets := splitLine[0][5:], strings.Split(splitLine[1], ";")

		for _, set := range sets {
			balls := strings.Split(set, ",")
			for _, ball := range balls {
				draw := strings.Split(strings.Trim(ball, " "), " ")
				number, color := utils.StrToInt(draw[0]), draw[1]

				if (color == "red" && number > 12) || (color == "green" && number > 13) || (color == "blue" && number > 14) {
					continue outer
				}
			}
		}

		sum += utils.StrToInt(game)
	}

	fmt.Println(sum)
}

func part2() {
	sum := 0

	for _, line := range data {
		splitLine := strings.Split(line, ":")
		sets := strings.Split(splitLine[1], ";")

		red, green, blue := 0, 0, 0
		for _, set := range sets {
			balls := strings.Split(set, ",")
			for _, ball := range balls {
				draw := strings.Split(strings.Trim(ball, " "), " ")
				number, color := utils.StrToInt(draw[0]), draw[1]

				if color == "red" && number > red {
					red = number
				}
				if color == "green" && number > green {
					green = number
				}
				if color == "blue" && number > blue {
					blue = number
				}
			}
		}

		sum += red * green * blue
	}

	fmt.Println(sum)
}
