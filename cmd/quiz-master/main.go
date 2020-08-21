package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const file = "problems.csv"

func main() {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bytes.NewReader(b))

	questions, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var correct int
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
			correct++
		}
	}

	fmt.Printf("You got %d out of %d questions correct!", correct, len(questions))
}
