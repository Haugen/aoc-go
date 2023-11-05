package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Haugen/aoc-go/internal/utils"
)

var posCount []int
var data []string

func main() {
	path, _ := os.Getwd()
	data = utils.FilenameToArray(path + "/input.txt")
	posCount = getPosCount()

	part1()
	part2()
}

func part1() {
	gamma := ""
	epsilon := ""

	for _, v := range posCount {
		if v > 0 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaDec, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonDec, _ := strconv.ParseInt(epsilon, 2, 64)
	result := gammaDec * epsilonDec

	fmt.Println(result)
}

func part2() {
	oxygen := findCriteria(data, 0, false)
	co2 := findCriteria(data, 0, true)

	gammaDec, _ := strconv.ParseInt(oxygen, 2, 64)
	epsilonDec, _ := strconv.ParseInt(co2, 2, 64)
	result := gammaDec * epsilonDec

	fmt.Println(result)
}

func getPosCount() []int {
	posCount := make([]int, 12)

	for _, line := range data {
		numbers := strings.Split(line, "")
		for i, num := range numbers {
			if n := utils.StrToInt(num); n == 1 {
				posCount[i]++
			} else {
				posCount[i]--
			}
		}
	}

	return posCount
}

func findCriteria(data []string, step int, reverse bool) string {
	newData := make([]string, 0)
	var one string
	var zero string

	if reverse {
		one = "0"
		zero = "1"
	} else {
		one = "1"
		zero = "0"
	}

	for _, line := range data {
		num := string(line[step])
		if posCount[step] >= 0 && num == one {
			newData = append(newData, line)
		} else if posCount[step] < 0 && num == zero {
			newData = append(newData, line)
		}
	}

	if len(newData) == 1 {
		return newData[0]
	}

	return findCriteria(newData, step+1, reverse)
}
