package main

import (
	"fmt"
	"os"

	"github.com/Haugen/aoc-go/internal/utils"
)

var data []string

func main2023() {
	path, _ := os.Getwd()
	data = utils.FilenameToArray(path + "/input.txt")

	part1()
}

func part1() {
	for _, line := range data {
		fmt.Println(line)
	}
}
