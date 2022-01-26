package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// scanning file name and taboo word from
	// command line
	var fileName, tabooWord string
	var isTaboo bool
	fmt.Scan(&fileName)
	fmt.Scan(&tabooWord)
	tabooWord = strings.ToLower(tabooWord)

	// opening file and checking existing test
	// or return error message, defer close function
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// assigning scanning function to the scanner variable
	scanner := bufio.NewScanner(file)

	// split each scanned line into words
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		if tabooWord == scanner.Text() {
			isTaboo = true
		}
	}
	fmt.Println(isTaboo)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// recommendations and hits
