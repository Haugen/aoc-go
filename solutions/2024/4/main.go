package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Haugen/aoc-go/internal/utils"
)

type worker struct {
	data [][]string
}

func main() {
	worker := Init()
	worker.MakeLines()
}

func Init() worker {
	path, _ := os.Getwd()
	var raw = utils.FilenameToArray(path + "/example.txt")

	var data [][]string
	for _, line := range raw {
		row := strings.Split(line, "")
		data = append(data, row)
	}

	return worker{
		data: data,
	}
}

func (w *worker) Find(row, col int) (int, bool) {
	c := w.data[row][col]
	check := []string{"M", "A", "S"}
	sum := 0

	if c == "X" {
		// Horizontal - CORRECT
		for i := 1; i <= 3; i++ {
			if col+i >= len(w.data[row]) {
				break
			}

			curr := w.data[row][col+i]
			if curr != check[i-1] {
				break
			}
			if i == 3 {
				sum++
			}
		}

		// Horizontal reverse - CORRECT
		for i := 1; i <= 3; i++ {
			if col-i < 0 {
				break
			}

			curr := w.data[row][col-i]
			if curr != check[i-1] {
				break
			}
			if i == 3 {
				sum++
			}
		}

		// Vertical - CORRECT
		for i := 1; i <= 3; i++ {
			if row+i >= len(w.data) {
				break
			}

			curr := w.data[row+i][col]
			if curr != check[i-1] {
				break
			}
			if i == 3 {
				sum++
			}
		}

		// Vertical reversed - CORRECT
		for i := 1; i <= 3; i++ {
			if row-i < 0 {
				break
			}

			curr := w.data[row-i][col]
			if curr != check[i-1] {
				break
			}
			if i == 3 {
				sum++
			}
		}

		// Diagonal 1 - CORRECT
		for i := 1; i <= 3; i++ {
			if row+i >= len(w.data) || col+i >= len(w.data[row]) {
				break
			}

			curr := w.data[row+i][col+i]
			if curr != check[i-1] {
				break
			}
			if i == 3 {
				sum++
			}
		}

		// Diagonal 1 reversed - CORRECT
		for i := 1; i <= 3; i++ {
			if row-i < 0 || col-i < 0 {
				break
			}

			curr := w.data[row-i][col-i]
			if curr != check[i-1] {
				break
			}
			if i == 3 {
				sum++
			}
		}

		// Diagonal 2
		for i := 1; i <= 3; i++ {
			if row+i >= len(w.data) || col-i < 0 {
				break
			}

			curr := w.data[row+i][col-i]
			if curr != check[i-1] {
				break
			}
			if i == 3 {
				sum++
			}
		}

		// Diagonal 2 reversed
		for i := 1; i <= 3; i++ {
			if row-i < 0 || col+i >= len(w.data[row]) {
				break
			}

			curr := w.data[row-i][col+i]
			if curr != check[i-1] {
				break
			}
			if i == 3 {
				sum++
			}
		}
	}

	return sum, sum > 0
}

func (w *worker) FindX(row, col int) bool {
	c := w.data[row][col]

	if row == 0 || row >= len(w.data)-1 || col == 0 || col >= len(w.data[row])-1 {
		return false
	}

	oneOk := false
	twoOk := false

	if c == "A" {
		topLeft := w.data[row-1][col-1]
		bottomLeft := w.data[row+1][col-1]
		topRight := w.data[row-1][col+1]
		bottomRight := w.data[row+1][col+1]

		if topLeft+bottomRight == "MS" || topLeft+bottomRight == "SM" {
			oneOk = true
		}
		if topRight+bottomLeft == "MS" || topRight+bottomLeft == "SM" {
			twoOk = true
		}
	}

	return oneOk && twoOk
}

func (w *worker) MakeLines() {
	part1 := 0
	part2 := 0
	for i := 0; i < len(w.data); i++ {
		for j := 0; j < len(w.data[i]); j++ {
			// Part 1
			// score, found := w.Find(i, j)
			// if found {
			// 	part1 += score
			// }

			// Part 2
			found := w.FindX(i, j)
			if found {
				part2++
			}
		}
	}
	fmt.Println(part1, part2)
}
