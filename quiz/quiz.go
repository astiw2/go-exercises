package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {

	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of questions,answer")
	timeLimit := flag.Int("limit", 15, "time limit for the quiz to complete in seconds")
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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, p := range problems {
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d. \n", correct, len(problems))
			return
		default:
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
	rand.Shuffle(len(problems), func(i, j int) {
		problems[i], problems[j] = problems[j], problems[i]
	})
	return problems
}
