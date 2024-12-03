package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/Haugen/aoc-go/internal/utils"
)

type worker struct {
	data     []string
	p1Result int
	p2Result int
}

func main() {
	worker := Init()
	worker.MakeLines()
}

func Init() worker {
	path, _ := os.Getwd()
	var data = utils.FilenameToArray(path + "/input.txt")

	return worker{
		data:     data,
		p1Result: 0,
		p2Result: 0,
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
	enabled := true
	for _, line := range w.data {
		splitLine := strings.Split(line, " ")

		for _, cl := range splitLine {
			for i := 0; i < len(cl)-6; i++ {
				mul := string(cl[i]) + string(cl[i+1]) + string(cl[i+2]) + string(cl[i+3])
				do := string(cl[i]) + string(cl[i+1]) + string(cl[i+2]) + string(cl[i+3])
				dont := string(cl[i]) + string(cl[i+1]) + string(cl[i+2]) + string(cl[i+3]) + string(cl[i+4]) + string(cl[i+5]) + string(cl[i+6])

				if do == "do()" {
					enabled = true
				}
				if dont == "don't()" {
					enabled = false
				}

				if mul == "mul(" {
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
					w.p1Result += result
					if enabled {
						w.p2Result += result
					}
				}
			}
		}
	}

	fmt.Println(w.p1Result, w.p2Result)
}
