package main

import (
	"fmt"
	"os"
	"slices"
	"unicode"

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
	var numbers [][]string

	for _, line := range data {
		var localNums []string
		for _, letter := range line {
			if unicode.IsNumber(letter) {
				localNums = append(localNums, string(letter))
			}
		}
		numbers = append(numbers, localNums)
	}

	total := 0
	for _, i := range numbers {
		first, last := i[0], i[len(i)-1]
		sum := utils.StrToInt(first + last)
		total += sum
	}

	fmt.Println(total)
}

func checkAppend(arr []string, test string) (bool, string) {
	if slices.Contains(arr, test) {
		return true, utils.IntToStr(slices.Index(arr, test) + 1)
	}
	return false, ""
}

func part2() {
	spelledNums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var numbers [][]string

	for _, line := range data {
		var localNums []string
		lineArr := utils.StrToArray(line)

		for i := 0; i < len(line); i++ {
			if unicode.IsNumber(rune(lineArr[i][0])) {
				localNums = append(localNums, string(line[i]))
			} else {
				if i < len(lineArr)-4 {
					test := lineArr[i] + lineArr[i+1] + lineArr[i+2] + lineArr[i+3] + lineArr[i+4]
					shouldAppend, value := checkAppend(spelledNums, test)
					if shouldAppend {
						localNums = append(localNums, value)
					}
				}
				if i < len(lineArr)-3 {
					test := lineArr[i] + lineArr[i+1] + lineArr[i+2] + lineArr[i+3]
					shouldAppend, value := checkAppend(spelledNums, test)
					if shouldAppend {
						localNums = append(localNums, value)
					}
				}
				if i < len(lineArr)-2 {
					test := lineArr[i] + lineArr[i+1] + lineArr[i+2]
					shouldAppend, value := checkAppend(spelledNums, test)
					if shouldAppend {
						localNums = append(localNums, value)
					}
				}
			}
		}

		numbers = append(numbers, localNums)
	}

	total := 0
	for _, i := range numbers {
		first, last := i[0], i[len(i)-1]
		sum := utils.StrToInt(first + last)
		total += sum
	}

	fmt.Println(total)
}
