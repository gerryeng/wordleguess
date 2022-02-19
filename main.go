package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var allWords []string

func main() {
	var correct [5]string
	var incorrect []string
	var presentCharsNotIn = map[string][]int{}

	readWords()
	fmt.Println(len(allWords), "words in list")

	var guess string
	var result string
	// Y = Correct character, correct place
	// X = Wrong character
	// P = Correct word, wrong position

	fmt.Print("Enter text: ")

	fmt.Scanln(&guess)
	fmt.Println("")
	fmt.Print("Enter result for ", guess, ": ")
	fmt.Scanln(&result)

	for pos, c := range result {
		char := string(c)
		if char == "Y" {
			correct[pos] = string(guess[pos])
		} else if char == "X" {
			incorrect = append(incorrect, string(guess[pos]))
		} else if char == "P" {
			presentCharsNotIn[string(guess[pos])] = append(presentCharsNotIn[string(guess[pos])], pos)
		}
	}

	fmt.Println("Correct ", correct)
	fmt.Println("Incorrect ", incorrect)
	fmt.Println("presentCharsNotIn ", presentCharsNotIn)

	// Print candidate words
	var candidates []string
	for _, w := range allWords {

		// check if string "w" contains any incorrect char
		containsIncorrectChar := false
		for _, c := range incorrect {
			if strContainsChar(w, c) {
				containsIncorrectChar = true
				break
			}
		}

		// check if string "w" contains correct chars but wrong position
		for presentChar, incorrectPositions := range presentCharsNotIn {
			fmt.Println(presentChar)
			fmt.Println(incorrectPositions)
			if strContainsChar(w, presentChar) {

			}
		}

		// Skip to next word
		if containsIncorrectChar {
			continue
		}
		for pos, c := range w {
			char := string(c)
			if char == correct[pos] {
				candidates = append(candidates, w)
			}
		}
	}
	fmt.Println(candidates)
}

func strContainsChar(str string, char string) int {
	for _, c := range str {
		if string(c) == char {
			return true
		}
	}
	return false
}

func readWords() {
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		allWords = append(allWords, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
