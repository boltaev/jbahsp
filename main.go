package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func scanTaboo(file *os.File) {
	var tabooWord string
	fmt.Scan(&tabooWord)
	tabooWord = strings.ToLower(tabooWord) // lowercase the string
	if tabooWord == "exit" {
		fmt.Println("Bye!")
		return
	}
	remarkTaboo(file, tabooWord)
}

func remarkTaboo(file *os.File, tabooWord string) {
	// assigning scanning function to the scanner variable
	scanner := bufio.NewScanner(file)
	// split each scanned line into words
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		if tabooWord == scanner.Text() {
			stars := len(tabooWord)
			tabooWord = ""
			for i := 0; i < stars; i++ {
				tabooWord += "*"
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(tabooWord)
	scanTaboo(file)
}

func main() {
	// scanning file name from command line
	var fileName string
	fmt.Scan(&fileName)
	// opening file and checking existing test
	// or return error message, defer close function
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanTaboo(file)
}

// recommendations and hits
