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
}

func part1() {
	dataArrStr := strings.Split(data[0], ",")
	dataArr := utils.StrArrayToIntArray(dataArrStr)

	var count int
	for i := 0; i < 80; i++ {
		count = 0
		for j := 0; j < len(dataArr); j++ {
			if dataArr[j] == 0 {
				count++
				dataArr[j] = 6
			} else {
				dataArr[j]--
			}
		}

		for x := 0; x < count; x++ {
			dataArr = append(dataArr, 8)
		}
	}

	fmt.Println(len(dataArr))
}
