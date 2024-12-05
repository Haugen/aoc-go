package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/Haugen/aoc-go/internal/utils"
)

type worker struct {
	rules   []string
	updates [][]string
}

func main() {
	worker := Init()
	worker.Check()
}

func Init() worker {
	path, _ := os.Getwd()
	raw := utils.FilenameToArray(path + "/input.txt")
	rules := make([]string, 0)
	updates := make([][]string, 0)
	makeUpdates := false

	for _, line := range raw {
		if makeUpdates {
			splitLine := strings.Split(line, ",")
			updates = append(updates, splitLine)
		} else {
			rules = append(rules, line)
			if line == "" {
				makeUpdates = true
				continue
			}
		}
	}

	return worker{
		rules:   rules,
		updates: updates,
	}
}

func (w *worker) Check() {
	p1 := 0
	for _, update := range w.updates {
		isGood, _ := w.CheckUpdate(update)
		if isGood {
			item := update[len(update)/2]
			p1 += utils.StrToInt(item)
		}
	}
	fmt.Println(p1)
}

func (w worker) CheckUpdate(update []string) (bool, int) {
	if len(update) <= 1 {
		return true, -1
	}

	first := update[0]
	for i := 1; i < len(update); i++ {
		test := update[i] + "|" + first
		found := slices.IndexFunc(w.rules, func(r string) bool { return r == test })

		if found != -1 {
			return false, i
		}
	}

	new := update[1:]
	return w.CheckUpdate(new)
}
