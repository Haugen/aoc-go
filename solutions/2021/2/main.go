package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Haugen/aoc-go/internal/utils"
)

func main() {
	path, _ := os.Getwd()
	data := utils.FilenameToArray(path + "/input.txt")

	part1(data)
	part2(data)
}

func part1(data []string) {
	hor := 0
	dep := 0

	for _, line := range data {
		lineArr := strings.Split(line, " ")
		a, b := lineArr[0], utils.StrToInt(lineArr[1])

		if a == "forward" {
			hor += b
		} else if a == "down" {
			dep += b
		} else if a == "up" {
			dep -= b
		}
	}

	fmt.Println(hor * dep)
}

func part2(data []string) {
	hor := 0
	dep := 0
	aim := 0

	for _, line := range data {
		lineArr := strings.Split(line, " ")
		a, b := lineArr[0], utils.StrToInt(lineArr[1])

		if a == "forward" {
			hor += b
			dep += aim * b
		} else if a == "down" {
			aim += b
		} else if a == "up" {
			aim -= b
		}
	}

	fmt.Println(hor * dep)
}
