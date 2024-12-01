package main

import (
	"fmt"
	"os"

	"github.com/Haugen/aoc-go/internal/utils"
)

var data []string

func main() {
	path, _ := os.Getwd()
	data = utils.FilenameToArray(path + "/example.txt")

	part1()
}

func part1() {
	lines := []string{}
	lines = append(lines, data...)

	for i := 0; i < len(lines); i++ {
		// partOf := false

		lineArr := utils.StrToArray(lines[i])

		fmt.Println(lineArr)

		for _, char := range lineArr {
			isInt := utils.IsInt(char)
		}

		if i == 0 {

		}

		if i > len(lines)-1 {

		}
	}
}
