package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

var client = &http.Client{}
var dayFlag = flag.Int("day", 0, "Select day")

func main() {
	flag.Parse()

	if *dayFlag == 0 {
		log.Fatal("You need to provide a day with the --day flag.")
	}
	if *dayFlag > 25 {
		log.Fatal("Please choose a day between 1 and 25")
	}

	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	binPath := filepath.Dir(ex)

	err = godotenv.Load(binPath + "/../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	year := os.Getenv("YEAR")
	dayString := strconv.FormatInt(int64(*dayFlag), 10)
	inputString := fetchInputData(year, dayString)

	folderPath := binPath + "/../" + year + "/" + dayString
	writeInputFile(folderPath, inputString)

	fmt.Printf("Successfully generated template and input file for day %s year %s. Good luck!\n", dayString, year)
}

func writeInputFile(folderPath string, input string) {
	os.MkdirAll(folderPath, os.ModePerm)

	f, err := os.Create(folderPath + "/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.WriteString(input)
}

func fetchInputData(year string, day string) string {
	aocSession := os.Getenv("AOC_SESSION")
	url := "https://adventofcode.com/" + year + "/day/" + day + "/input"

	req, err := http.NewRequest("GET", url, nil)
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
