package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/Haugen/aoc-go/internal/utils"
)

type worker struct {
	data []string
}

func main() {
	worker := Init()
	worker.MakeLines()
}

func Init() worker {
	path, _ := os.Getwd()
	var data = utils.FilenameToArray(path + "/input.txt")

	return worker{
		data: data,
	}
}

func (w worker) ProduceNumList(cl string, i int) ([]string, int) {
	nums := make([]string, 0)
	for {
		isInt := unicode.IsDigit(rune(cl[i]))
		if !isInt {
			break
		}

		if isInt {
			nums = append(nums, string(cl[i]))
			i++
		}
	}

	return nums, i
}

func (w worker) Result(leftNums []string, rightNums []string) int {
	var leftNum string
	for _, num := range leftNums {
		leftNum += num
	}
	var rightNum string
	for _, num := range rightNums {
		rightNum += num
	}

	result := utils.StrToInt(leftNum) * utils.StrToInt(rightNum)
	return result
}

func (w *worker) MakeLines() {
	sum := 0

	for _, line := range w.data {
		splitLine := strings.Split(line, " ")

		for _, cl := range splitLine {
			for i := 0; i < len(cl)-4; i++ {
				word := string(cl[i]) + string(cl[i+1]) + string(cl[i+2]) + string(cl[i+3])
				if word == "mul(" {
					first := i + 4

					leftNums, k := w.ProduceNumList(cl, first)

					if len(leftNums) == 0 || string(cl[k]) != string(',') {
						continue
					}

					rightNums, k := w.ProduceNumList(cl, k+1)

					if len(rightNums) == 0 || string(cl[k]) != string(')') {
						continue
					}

					result := w.Result(leftNums, rightNums)
					sum += result
				}
			}
		}
	}

	fmt.Println(sum)
}
