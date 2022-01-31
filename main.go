package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func scanTaboo() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	inputArr := strings.Split(line, " ")
	return inputArr
}

func remarkTaboo(badWords map[int]string, tabooWord []string) {
	var result string
	// comparing words from two arrays
	for i := 0; i < len(badWords); i++ {
		for j := 0; j < len(tabooWord); j++ {
			if badWords[i] == strings.ToLower(tabooWord[j]) {
				stars := len(tabooWord[j])
				tabooWord[j] = ""
				for p := 0; p < stars; p++ {
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

	badWords := make(map[int]string)
	var tabooWord []string
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

	var i int
	for scanner.Scan() {
		badWords[i] = scanner.Text()
		i++
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
