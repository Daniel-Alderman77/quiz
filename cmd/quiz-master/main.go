package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

const (
	file        = "problems.csv"
	defaultTime = time.Second * 30
)

func main() {
	timeLimit := flag.Duration("defaultTime", defaultTime, "a duration, i.e: 30s")
	flag.Parse()

	questions := parseFile()

	fmt.Printf("You have %v to answer as many questions as possible, press enter to start!", timeLimit)

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	correct := make(chan bool, 1)
	go func() {
		correct <- askQuestion(questions)
	}()

	select {
	case <-time.After(*timeLimit):
		fmt.Printf("Out of time! You got %d out of %d file correct!\n", len(correct), len(questions))
	}

	fmt.Printf("You got %d out of %d questions correct!\n", len(correct), len(questions))
}

func parseFile() [][]string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bytes.NewReader(b))

	questions, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return questions
}

func askQuestion(questions [][]string) bool {
	for i, q := range questions {
		fmt.Printf("Q%d, what is %v?\n", i+1, q[0])

		reader := bufio.NewReader(os.Stdin)

		answer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		// convert CRLF to LF
		answer = strings.Replace(answer, "\n", "", -1)

		if answer == q[1] {
			return true
		}

		return false
	}

	return false
}
