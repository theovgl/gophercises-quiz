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

func printHelp() {
	fmt.Println("Usage of ./quiz:")

	fmt.Println("  -h")
	fmt.Println("\tDisplay this help message")

	fmt.Println("  -csv string")
	fmt.Println("\tA csv file in the format 'question,answer' (default \"problems.csv\")")

	fmt.Println("  -limit int")
	fmt.Println("\tthe time limit for the quiz in seconds (default 30)")
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
	helpFlag := flag.Bool("h", false, "display help")
	csvFlag := flag.String("csv", "problems.csv", "specify file path")
	timeFlag := flag.Int("limit", 30, "specify time limit")

	// Parse flags
	flag.Parse()

	if *helpFlag {
		printHelp()
		os.Exit(0)
	}

	problems, err := parseCSVFile(*csvFlag)
	checkError(err)

	config := &QuizConfig{
		timeLimit: *timeFlag,
		score:     0,
	}

	runQuiz(config, problems)
}
