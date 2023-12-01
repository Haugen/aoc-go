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
					if slices.Contains(spelledNums, test) {
						localNums = append(localNums, utils.IntToStr(slices.Index(spelledNums, test)+1))
					}
				}
				if i < len(lineArr)-3 {
					test := lineArr[i] + lineArr[i+1] + lineArr[i+2] + lineArr[i+3]
					if slices.Contains(spelledNums, test) {
						localNums = append(localNums, utils.IntToStr(slices.Index(spelledNums, test)+1))
					}
				}
				if i < len(lineArr)-2 {
					test := lineArr[i] + lineArr[i+1] + lineArr[i+2]
					if slices.Contains(spelledNums, test) {
						localNums = append(localNums, utils.IntToStr(slices.Index(spelledNums, test)+1))
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
