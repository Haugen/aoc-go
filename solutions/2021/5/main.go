package main

import (
	"fmt"
	"os"
	"strings"

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
	plane := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		plane[i] = make([]int, 1000)
	}

	for _, line := range data {
		coords := strings.Split(line, " -> ")
		first, last := coords[0], coords[1]
		x := strings.Split(first, ",")
		y := strings.Split(last, ",")
		x1, y1 := utils.StrToInt(x[0]), utils.StrToInt(x[1])
		x2, y2 := utils.StrToInt(y[0]), utils.StrToInt(y[1])

		if x1 == x2 {
			if y1 < y2 {
				for i := y1; i <= y2; i++ {
					plane[i][x1]++
				}
			} else {
				for i := y1; i >= y2; i-- {
					plane[i][x1]++
				}
			}
		} else if y1 == y2 {
			if x1 < x2 {
				for i := x1; i <= x2; i++ {
					plane[y1][i]++
				}
			} else {
				for i := x1; i >= x2; i-- {
					plane[y1][i]++
				}
			}
		}
	}

	num := 0
	for _, row := range plane {
		for _, cell := range row {
			if cell > 1 {
				num++
			}
		}
	}

	fmt.Println(num)
}

func part2() {
	plane := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		plane[i] = make([]int, 1000)
	}

	for _, line := range data {
		coords := strings.Split(line, " -> ")
		first, last := coords[0], coords[1]
		x := strings.Split(first, ",")
		y := strings.Split(last, ",")
		x1, y1 := utils.StrToInt(x[0]), utils.StrToInt(x[1])
		x2, y2 := utils.StrToInt(y[0]), utils.StrToInt(y[1])

		if x1 == x2 {
			if y1 < y2 {
				for i := y1; i <= y2; i++ {
					plane[i][x1]++
				}
			} else {
				for i := y1; i >= y2; i-- {
					plane[i][x1]++
				}
			}
		} else if y1 == y2 {
			if x1 < x2 {
				for i := x1; i <= x2; i++ {
					plane[y1][i]++
				}
			} else {
				for i := x1; i >= x2; i-- {
					plane[y1][i]++
				}
			}
			// Part 2 addition
		} else {
			if x1 > x2 && y1 > y2 {
				for i := 0; i <= x1-x2; i++ {
					plane[y1-i][x1-i]++
				}
			} else if x1 > x2 && y1 < y2 {
				for i := 0; i <= x1-x2; i++ {
					plane[y1+i][x1-i]++
				}
			} else if x1 < x2 && y1 > y2 {
				for i := 0; i <= x2-x1; i++ {
					plane[y1-i][x1+i]++
				}
			} else if x1 < x2 && y1 < y2 {
				for i := 0; i <= x2-x1; i++ {
					plane[y1+i][x1+i]++
				}
			}
		}
	}

	num := 0
	for _, row := range plane {
		for _, cell := range row {
			if cell > 1 {
				num++
			}
		}
	}

	fmt.Println(num)
}

// Used to visualize the plane
func printGrid(plane [][]int) {
	for _, row := range plane {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(cell)
			}
		}
		fmt.Println()
	}
}
