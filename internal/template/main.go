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
}

func part1(data []string) {
	for _, line := range data {
		fmt.Println(line)
	}
}
