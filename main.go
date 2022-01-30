package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func scanTaboo() []string {
	var tabooWord string
	fmt.Scan(&tabooWord)
	inputArr := strings.Split(tabooWord, " ")
	return inputArr
}

func remarkTaboo(badWords []string, tabooWord []string) {
	var result string
	inputArr := tabooWord // here assigning new variable for case
	// lowering words
	for _, s := range inputArr {
		s = strings.ToLower(s)
	}
	// comparing words from two arrays
	for i := 0; i < len(badWords); i++ {
		for j := 0; j < len(inputArr); j++ {
			if badWords[i] == inputArr[j] {
				stars := len(inputArr[j])
				tabooWord[j] = ""
				for i := 0; i < stars; i++ {
					tabooWord[j] += "*"
				}
			}
		}
	}
	// uniting words from arrays to the string
	for q := 0; q < len(tabooWord); q++ {
		result += tabooWord[q] + " "
	}

	fmt.Println(result)
}

func main() {
	// scanning file name from command line
	var fileName string
	var badWords, tabooWord []string

	fmt.Scan(&fileName)
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

	for i := 0; i < len(scanner.Text()); i++ {
		badWords[i] = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for {
		tabooWord = scanTaboo()
		if tabooWord[0] == "exit" {
			fmt.Println("Bye!")
			break
		}
		remarkTaboo(badWords, tabooWord)
	}
}
