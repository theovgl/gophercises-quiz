package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Problem struct {
	question string
	answer   string
}

type QuizConfig struct {
	timeLimit int
	score     int
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func runQuiz(config *QuizConfig, problems []Problem) {
	for i, problem := range problems {
		var answer string
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)

		if _, err := fmt.Scanln(&answer); err != nil && err.Error() != "unexpected newline" {
			fmt.Println("Error reading input:", err)
			continue
		}

		if answer == problem.answer {
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

func parseProblems(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))

	for i, line := range lines {
		problem := Problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
		problems[i] = problem
	}

	return problems
}

func main() {
	// Declare flags
	csvFlag := flag.String("csv", "problems/maths.csv", "a csv file in the format 'question,answer'")
	timeFlag := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	// Parse flags
	flag.Parse()

	lines, err := parseCSVFile(*csvFlag)
	checkError(err)

	problems := parseProblems(lines)

	config := &QuizConfig{
		timeLimit: *timeFlag,
		score:     0,
	}

	runQuiz(config, problems)
}
