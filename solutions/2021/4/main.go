package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Haugen/aoc-go/internal/utils"
)

var data []string

type boardItem struct {
	number string
	drawn  bool
}

type board [][]boardItem

type boardType struct {
	board board
	won   bool
}

func main() {
	path, _ := os.Getwd()
	data = utils.FilenameToArray(path + "/input.txt")

	part1()
	part2()
}

func part2() {
	numbers, rest := strings.Split(data[0], ","), data[2:]
	boards := buildBoards(rest)
	winners := 0

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(boards); j++ {
			num := numbers[i]
			doBoard(boards[j], num)
			if !boards[j].won && checkWinner(boards[j].board) {
				winners++
				boards[j].won = true
				// Only works if all boards eventually win
				if winners == len((boards)) {
					calculateWinner(boards[j], num)
				}
			}
		}
	}
}

func part1() {
	numbers, rest := strings.Split(data[0], ","), data[2:]
	boards := buildBoards(rest)

outer:
	for _, num := range numbers {
		for _, board := range boards {
			doBoard(board, num)
			if checkWinner(board.board) {
				calculateWinner(board, num)
				break outer
			}
		}
	}
}

func doBoard(board boardType, num string) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board.board[i][j].number == num {
				board.board[i][j].drawn = true
			}
		}
	}
}

func checkWinner(board board) bool {
	winner := false

	// Rows
outerRow:
	for _, row := range board {
		for i := 0; i < 5; i++ {
			if !row[i].drawn {
				break
			}
			if i == 4 {
				winner = true
				break outerRow
			}
		}
	}

	// Cols
	if !winner {
	outerCol:
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if !board[j][i].drawn {
					break
				}
				if j == 4 {
					winner = true
					break outerCol
				}
			}
		}
	}

	return winner
}

func calculateWinner(board boardType, num string) {
	var unmarked []string

	for _, row := range board.board {
		for _, item := range row {
			if !item.drawn {
				unmarked = append(unmarked, item.number)
			}
		}
	}

	unmarkedInt := utils.StrArrayToIntArray(unmarked)
	sum := utils.SumArray(unmarkedInt)

	// Result
	fmt.Println(sum * utils.StrToInt(num))
}

func buildBoards(lines []string) []boardType {
	boards := []boardType{}

	for i := 0; i < len(lines); i += 6 {
		board := board{}

		for j := 0; j < 5; j++ {
			row := []boardItem{}
			for _, num := range strings.Fields(lines[i+j]) {
				row = append(row, boardItem{number: num, drawn: false})
			}
			board = append(board, row)
		}
		boards = append(boards, boardType{board: board, won: false})
	}

	return boards
}
