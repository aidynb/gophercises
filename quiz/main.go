package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	parseArgs(os.Args[1:])

	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalf("Error opening file: %s\n", err.Error())
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Error reading file: %s\n", err.Error())
	}

	buf := bufio.NewReader(os.Stdin)

	var (
		totalQuestions int = len(records)
		counter        int
	)

	for i := range records {
		question, answer := records[i][0], records[i][1]

		fmt.Printf("Question %d: %s > ", i+1, question)
		input, err := buf.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading from stdin: %s\n", err.Error())
		}
		input = strings.TrimRight(input, "\n")

		if input == answer {
			counter++
		}
	}

	fmt.Printf("You've got %v out of %v\n", counter, totalQuestions)
}

// TODO
func parseArgs(args []string) {
	for i := range args {
		if args[i] == "-h" || args[i] == "--help" {
			showHelp()
		}
	}
}

// TODO
func showHelp() {
	fmt.Printf("Usage:")
}
