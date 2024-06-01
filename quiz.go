package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type QuizConfig struct {
	timeLimit int
	score     int
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func runQuiz(config *QuizConfig, problems [][]string) {
	for i, problem := range problems {
		var answer string
		fmt.Printf("Problem #%d: %s = ", i+1, problem[0])

		if _, err := fmt.Scanln(&answer); err != nil && err.Error() != "unexpected newline" {
			fmt.Println("Error reading input:", err)
			continue
		}

		if answer == problem[1] {
			config.score++
		}
	}
	fmt.Printf("You scored %d out of %d\n", config.score, len(problems))
}

func parseCSVFile(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func main() {
	// Declare flags
	csvFlag := flag.String("csv", "problems/maths.csv", "a csv file in the format 'question,answer'")
	timeFlag := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	// Parse flags
	flag.Parse()

	problems, err := parseCSVFile(*csvFlag)
	checkError(err)

	config := &QuizConfig{
		timeLimit: *timeFlag,
		score:     0,
	}

	runQuiz(config, problems)
}
