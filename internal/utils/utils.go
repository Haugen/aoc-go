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
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func StrArrayToIntArray(arr []string) []int {
	var newArr []int
	for _, v := range arr {
		newArr = append(newArr, StrToInt(v))
	}
	return newArr
}

func SumArray(array []int) int {
	sum := 0
	for _, v := range array {
		sum += v
	}
	return sum
}
