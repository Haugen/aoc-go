package main

import (
	"fmt"
	"math"
	"os"
	"strings"

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

func isGood(levels []string) bool {
	first := utils.StrToInt(levels[0])
	second := utils.StrToInt(levels[1])

	if first == second {
		return false
	}

	if first > second {
		for i := 0; i < len(levels)-1; i++ {
			curr := utils.StrToInt(levels[i])
			next := utils.StrToInt(levels[i+1])
			diff := math.Abs(float64(curr) - float64(next))

			if curr <= next || diff > 3 {
				return false
			}
		}
		return true
	}

	if first < second {
		for i := 0; i < len(levels)-1; i++ {
			curr := utils.StrToInt(levels[i])
			next := utils.StrToInt(levels[i+1])
			diff := math.Abs(float64(curr) - float64(next))

			if curr >= next || diff > 3 {
				return false
			}
		}
		return true
	}

	return false
}

func (w *worker) MakeLines() {
	p1 := 0
	p2 := 0

Loop:
	for _, report := range w.data {
		levels := strings.Split(report, " ")

		if isGood(levels) {
			p1++
			p2++
			continue
		}

		for i := 0; i < len(levels); i++ {
			new := make([]string, len(levels))
			copy(new, levels)
			new = append(new[:i], new[i+1:]...)
			if isGood(new) {
				p2++
				continue Loop
			}
		}

	}

	fmt.Println(p1, p2)
}
