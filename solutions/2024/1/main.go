package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/Haugen/aoc-go/internal/utils"
)

type worker struct {
	data     []string
	left     []int
	right    []int
	rightOcc map[int]int
}

func main() {
	worker := Init()
	worker.MakeLines()
	worker.SortLines()
	part1 := worker.SumDistance()
	fmt.Println(part1)

	worker.BuildMap()
	part2 := worker.SumSimilarities()
	fmt.Println(part2)
}

func Init() worker {
	path, _ := os.Getwd()
	data := utils.FilenameToArray(path + "/input.txt")

	return worker{
		data:     data,
		left:     make([]int, 0),
		right:    make([]int, 0),
		rightOcc: make(map[int]int),
	}
}

func (w *worker) MakeLines() {
	for _, line := range w.data {
		splitLine := strings.Split(line, "   ")
		w.left = append(w.left, utils.StrToInt(splitLine[0]))
		w.right = append(w.right, utils.StrToInt(splitLine[1]))
	}
}

func (w *worker) SortLines() {
	sort.Slice(w.left, func(i, j int) bool {
		return w.left[i] < w.left[j]
	})
	sort.Slice(w.right, func(i, j int) bool {
		return w.right[i] < w.right[j]
	})

}

func (w *worker) SumDistance() int {
	sum := 0
	for i, leftItem := range w.left {
		sum += int(math.Abs(float64(leftItem - w.right[i])))
	}
	return sum
}

func (w *worker) BuildMap() {
	for _, line := range w.right {
		_, ok := w.rightOcc[line]

		if ok {
			w.rightOcc[line]++
		} else {
			w.rightOcc[line] = 1
		}
	}
}

func (w *worker) SumSimilarities() int {
	sum := 0

	for _, line := range w.left {
		val, ok := w.rightOcc[line]

		if ok {
			sum += line * val
		}
	}

	return sum
}
