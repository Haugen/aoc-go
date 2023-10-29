package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func FilenameToArray(filename string) []string {
	input, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	var inputLines []string

	for fileScanner.Scan() {
		inputLines = append(inputLines, fileScanner.Text())
	}

	return inputLines
}

func StrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
