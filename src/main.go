package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var client = &http.Client{}

func main() {
	var dayFlag = flag.Int("day", 0, "Select day")
	flag.Parse()

	if *dayFlag == 0 {
		log.Fatal("You need to provide a day with the --day flag.")
	}
	if *dayFlag > 25 {
		log.Fatal("Please choose a day between 1 and 25")
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	year := os.Getenv("YEAR")
	inputString := fetchInputData(year, strconv.FormatInt(int64(*dayFlag), 10))
	fmt.Println(inputString)
}

func fetchInputData(year string, day string) string {
	aocSession := os.Getenv("AOC_SESSION")
	req, err := http.NewRequest("GET", "https://adventofcode.com/"+year+"/day/"+day+"/input", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("cookie", "session="+aocSession)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
