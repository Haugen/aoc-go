package main

import (
	"fmt"
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

func (w *worker) MakeLines() {
	for _, line := range w.data {
		splitLine := strings.Split(line, " ")
		fmt.Println(splitLine)
	}
}
