package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// Keep track of right/wrong answers
	correctAnswers := 0
	// Flag creation
	fileNamePtr := flag.String("f", "problems.csv", "Name of file with quiz questions")
	flag.Parse()

	// Read in CSV and work with data
	file, err := os.Open(*fileNamePtr)
	if err != nil {
		fmt.Println("Error:", err)
	}
	r := csv.NewReader(file)
	data, err := r.ReadAll()
	timer := time.NewTimer(5 * time.Second)
	for {
		go runQuiz(data, &correctAnswers)
		<-timer.C
		break
	}

	// Print total correct answers and total questions asked
	fmt.Printf("\nYou answered %d questions correctly out of %d", correctAnswers, len(data))

}

func runQuiz(data [][]string, correctAnswers *int) {
	for _, line := range data {
		q := line[0]
		a := line[1]
		fmt.Printf("What is %s?\t", q)
		userAnswer := ""
		fmt.Scanln(&userAnswer)
		if userAnswer == a {
			*correctAnswers += 1
		}
	}
}
