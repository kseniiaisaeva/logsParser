package main

import (
	"bufio"
	"fmt"
	"github.com/kseniiaisaeva/parsing-logs/parser"
	"log"
	"os"
)

var urls = make(map[string]int)
var ips = make(map[string]int)

const format = "The number of unique IP addresses: %d\nThe top 3 most visited URLs: %v\nThe top 3 most active IP addresses: %v"

func main() {
	var filePath string
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	} else {
		filePath = "./programming-task/programming-task-example-data.log"
	}
	readInput(filePath)

	writeOutput()
}

// Read string and extracts info and update maps
func processLogLine(logLine string) {
	ip, url, err := parser.ExtractInfo(logLine)
	if err != nil {
		fmt.Printf("Unable to read line %s: %v", logLine, err)
		return
	}
	ips = parser.AddNewEntry(ips, ip)
	urls = parser.AddNewEntry(urls, url)
}

// Write results to output file
func writeOutput() {
	uniqueIps, topUrls, topIps := getResults()
	res := fmt.Sprintf(format, uniqueIps, topUrls, topIps)
	fmt.Println(res)
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(res)
	if err != nil {
		fmt.Printf("Unable to write output: %v", err)
	}
}

// Get results: number of unique ips, top 3 visited urls and top 3 active ips
func getResults() (int, []string, []string) {
	return len(ips), parser.GetTopThree(urls), parser.GetTopThree(ips)
}

func readInput(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		processLogLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
