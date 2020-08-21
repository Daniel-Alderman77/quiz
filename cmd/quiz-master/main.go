package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
)

const file = "problems.csv"

func main() {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bytes.NewReader(b))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)
}
