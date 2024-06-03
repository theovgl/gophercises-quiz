package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Problem struct {
	question string
	answer   string
}

type QuizConfig struct {
	timeLimit int
	score     int
}

func printScore(outOfTime bool, config *QuizConfig, problemsNumber int) {
	if outOfTime {
		fmt.Printf("\nTime's up ! You scored %d out of %d\n", config.score, problemsNumber)
	} else {
		fmt.Printf("\nYou scored %d out of %d\n", config.score, problemsNumber)
	}
}

func runQuiz(config *QuizConfig, problems []Problem) {
	// Start Timer
	timer := time.NewTimer(time.Duration(config.timeLimit) * time.Second)
	defer timer.Stop()

	done := make(chan bool)

	// Start quizz loop
	go func() {
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
		done <- true
	}()

	select {
	case <-timer.C:
		printScore(true, config, len(problems))
	case <-done:
		printScore(false, config, len(problems))
	}
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

func parseProblems(lines [][]string) ([]Problem, error) {
	problems := make([]Problem, len(lines))

	for i, line := range lines {
		if len(line) != 2 {
			err := errors.New("wrong CSV format")
			return nil, err
		}

		problem := Problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
		problems[i] = problem
	}

	return problems, nil
}

func main() {
	// Declare flags
	csvFlag := flag.String("csv", "problems/maths.csv", "a csv file in the format 'question,answer'")
	timeFlag := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	// Parse flags
	flag.Parse()

	lines, err := parseCSVFile(*csvFlag)
	if err != nil {
		log.Fatal(err)
	}

	problems, err := parseProblems(lines)
	if err != nil {
		log.Fatal(err)
	}

	config := &QuizConfig{
		timeLimit: *timeFlag,
		score:     0,
	}

	runQuiz(config, problems)
}
