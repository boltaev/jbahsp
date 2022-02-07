package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func scanTaboo() []string {
	fmt.Print("Please enter your text(or enter exit for stop chatting): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	introLine := strings.Split(line, " ")
	return introLine
}

func remarkTaboo(tabooWords map[int]string, introLine []string) {
	var result string
	// comparing words from two arrays
	for i := 0; i < len(tabooWords); i++ {
		for j := 0; j < len(introLine); j++ {
			if tabooWords[i] == strings.ToLower(introLine[j]) {
				stars := len(introLine[j])
				introLine[j] = ""
				for p := 0; p < stars; p++ {
					introLine[j] += "*"
				}
			}
		}
	}
	// uniting words from arrays to the string
	for q := 0; q < len(introLine); q++ {
		result += introLine[q] + " "
	}
	fmt.Print("Checked text: ")
	fmt.Println(result)
}

func mainLoop(tabooWords map[int]string, introLine []string) {
	for {
		introLine = scanTaboo()
		if introLine[0] == "exit" {
			fmt.Println("Bye!")
			break
		}
		remarkTaboo(tabooWords, introLine)
	}
}

func main() {
	// scanning file name from command line
	var fileName string

	tabooWords := make(map[int]string)
	var introLine []string
	fmt.Print("Please enter name of txt format file: ")
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
		tabooWords[i] = scanner.Text()
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	mainLoop(tabooWords, introLine)
}

// comment section for notes
// when we enter second command on CL
// loop works without reason as an empty function

// future plans for this project:
// commenting every function and action
// for more details
