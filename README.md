<h1 align="center">
:christmas_tree: Advent of Code in Go :christmas_tree:
</h1>

```go
go run main.go --day=1
```

aoc-go is a repository intended to automate some tasks related to Advent of Code, like fetching input data and provide a starting template for each day, as well as providing utility functions for commonly used tasks. The repository should be forked and used to solve puzzles directly in it. If a `solutions` folder is present, delete it after forking to get started on your own.

## Get your session token

First thing you need is your session token from adventofcode.com. This is used to fetch your personalized input data from each puzzle. It can be found under the network tab in your dev tools. [Check this post for assistance](https://github.com/wimglenn/advent-of-code-wim/issues/1).

Once you have your token, copy .env.example, rename it to .env, and include your token in the AOC_SESSION spot. In your .env file you can also adjust what year of Advent of Code you want to work with. Defaults to 2022.

## Usage

- Make sure you have your `.env` file setup and both the `AOC_SESSION` and `YEAR` variables set.
- From the root of this project, run for example `go run main.go --day=1`. This will fetch data for day 1 for the configured year and set up a template for you to get started.
- Templates will be created at the path `solutions/{year}/{day}`.

## The template

The template to get started each day looks like below. Additional utilities are provided in the `utils` package. Template has been updated in 2024 to practice more of a "Go style", working with a single struct, adding datastructures and methods to it.

```go
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

```

## Useful utils while solving puzzles

All utility functions are provided in the github.com/Haugen/aoc-go/internal/utils package.

| Function                   | Description                                   |
| -------------------------- | --------------------------------------------- |
| `utils.IntToStr`           | Convert an int to a string                    |
| `utils.IsInt`              | Check if string is an int                     |
| `utils.StrToInt`           | Convert a string to an int                    |
| `utils.StrToArray`         | Convert a string to an array of strings       |
| `utils.StrArrayToIntArray` | Convert an array of string to an array of int |
| `utils.SumArray`           | Sum array of int                              |
