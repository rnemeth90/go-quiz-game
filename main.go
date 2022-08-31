package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	var correct int
	scanner := bufio.NewScanner(os.Stdin)
	fileName := "questions.csv"
	csvFile, err := os.Open(fileName)
	check(err)

	reader := csv.NewReader(csvFile)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		check(err)
		fmt.Print(record[0] + " ")
		scanner.Scan()
		a := scanner.Text()
		num1, e := strconv.Atoi(record[1])
		check(e)
		num2, e := strconv.Atoi(a)
		check(e)

		if num1 == num2 {
			correct++
		}
	}

	fmt.Println("")
	fmt.Printf("Total correct: %d\n", correct)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
