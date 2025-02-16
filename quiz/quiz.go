package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	question string
	answer   string
}

func main() {

	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of questions,answer")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Printf("Failed to open the file: %s\n", *csvFileName)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Failed to parse the csv")
	}
	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem no. %d: %s = \n", i+1, p.question)
		var answer string
		_, err = fmt.Scanf("%s\n", &answer)
		if err != nil {
			fmt.Println("Failed to read the answer", err)
		}
		if answer == p.answer {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d. \n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return problems
}
