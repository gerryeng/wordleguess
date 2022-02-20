package main

import (
	"bufio"
	"log"
	"os"
)

// Helper script to filter 5 letter words from list of possible english words

// What it does:
// Reads allwords.txt, containing all possible English words
// select only words with 5 characters and saves it in words.txt
//
// output from words.txt should then be saved in words/words.go
func main() {
	file, err := os.Open("allwords.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	output, err := os.OpenFile("words.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) == 5 {
			if _, err := output.WriteString(word + "\n"); err != nil {
				panic(err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
