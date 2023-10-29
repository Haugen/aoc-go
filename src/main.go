package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var client = &http.Client{}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AOC_SESSION := os.Getenv("AOC_SESSION")
	YEAR := os.Getenv("YEAR")

	fmt.Println(AOC_SESSION)
	fmt.Println(YEAR)

	req, err := http.NewRequest("GET", "https://adventofcode.com/"+YEAR+"/day/1/input", nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}
	req.Header.Add("cookie", "session="+AOC_SESSION)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response body:", err)
	}

	fmt.Println(string(body))
}
