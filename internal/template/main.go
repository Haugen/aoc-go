package main

import (
	"fmt"
	"os"

	"github.com/Haugen/aoc-go/internal/utils"
)

func main() {
	path, _ := os.Getwd()
	data := utils.FilenameToArray(path + "/input.txt")

	for _, line := range data {
		fmt.Println(line)
	}
}
