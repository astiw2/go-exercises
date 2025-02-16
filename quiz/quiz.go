package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of questions,answer")
	flag.Parse()

	_, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Printf("Failed to read the file: %s", *csvFileName)
	}
}
