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
	worker.Do()
}

func Init() worker {
	path, _ := os.Getwd()
	raw := utils.FilenameToArray(path + "/input.txt")
	data := make([][]string, len(raw))

	for _, line := range raw {
		splitLine := strings.Split(line, "")
		data = append(data, splitLine)
	}

	return worker{
		data: data,
	}
}

func (w *worker) Do() {
	fmt.Println(w.data)
}
