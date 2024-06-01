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
	fmt.Println("Usage of ./quizz:")

	fmt.Println("  -h")
	fmt.Println("\tDisplay this help message")

	fmt.Println("  -csv string")
	fmt.Println("\tA csv file in the format 'question,answer' (default \"problems.csv\")")

	fmt.Println("  -limit int")
	fmt.Println("\tthe time limit for the quizz in seconds (default 30)")
}

func checkError(e error) {
	if e != nil {
		panic(e)
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
}
