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
	dayFlag := flag.Int("day", 0, "Select day")
	flag.Parse()

	if *dayFlag == 0 {
		log.Fatal("You need to provide a day with the --day flag.")
	}
	if *dayFlag > 25 {
		log.Fatal("Please choose a day between 1 and 25")
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	err = godotenv.Load(path + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	year := os.Getenv("YEAR")
	dayString := strconv.FormatInt(int64(*dayFlag), 10)
	folderPath := path + "/solutions/" + year + "/" + dayString

	inputString := fetchInputData(year, dayString)
	writeInputFile(folderPath, inputString)
	copyTemplateFile(path, folderPath)

	fmt.Printf("Successfully generated template and input file for day %s year %s. Good luck!\n", dayString, year)
}

func copyTemplateFile(binPath string, folderPath string) int64 {
	src := binPath + "/internal/template/main.go"

	source, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer source.Close()

	dest, err := os.Create(folderPath + "/main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()

	nBytes, err := io.Copy(dest, source)
	if err != nil {
		log.Fatal(err)
	}

	return nBytes
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
