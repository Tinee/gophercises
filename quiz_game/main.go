package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	var (
		filePath = flag.String("path", "problems.csv", "path to csv file, default to problems.csv")
		duration = flag.Int("timer", 30, "timer until it's game over. (defaults to 30)")
	)
	flag.Parse()

	f, err := os.Open(*filePath)
	timer := time.NewTimer(time.Second * time.Duration(*duration))

	if err != nil {
		log.Fatalln("Failed to open file")
	}

	lines, err := csv.NewReader(f).ReadAll()

	if err != nil {
		log.Fatalln("Failed to read csv file")
	}

	score := 0

problemLoop:
	for _, p := range parseProblems(lines) {
		answerCh := make(chan string)
		fmt.Printf("What's %v?", p.q)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			break problemLoop
		case answer := <-answerCh:
			if p.a == answer {
				score++
				continue problemLoop
			}
		}

	}

	fmt.Printf("Your score ended on %v out of %v", score, len(lines))
}

func parseProblems(lines [][]string) (problems []problem) {
	for _, l := range lines {
		problems = append(problems, problem{
			q: l[0],
			a: l[1],
		})
	}

	return problems
}
