package main

import (
	"fmt"

	"github.com/gerryeng/wordleguess/words"
)

func main() {
	var correct [5]string
	var incorrectChars []string
	var correctCharsWrongPositions = map[string][]int{}

	possibleWords := words.WORDS
	fmt.Println(len(possibleWords), "possible words in list")

	var guess string
	var result string
	for {

		fmt.Print("Enter guess: ")
		fmt.Scanln(&guess)
		fmt.Println("Enter result for guess")
		fmt.Println("\tY = Correct letter correct position")
		fmt.Println("\tX = Incorrect letter")
		fmt.Println("\tP = Correct letter, incorrect position")
		fmt.Println("\tExample 1: all letters are correct, enter YYYYY")
		fmt.Println("\tExample 2: 1st letter correct, 2nd letter is incorrect position, enter YPXXX")
		fmt.Print("Enter result for your guess (e.g. YYYYY, XXXXX, YPXXX):")
		fmt.Scanln(&result)

		// Y = Correct character, correct place
		// X = Wrong character
		// P = Correct word, wrong position
		for pos, c := range result {
			char := string(c)
			if char == "Y" {
				correct[pos] = string(guess[pos])
			} else if char == "X" {
				incorrectChars = append(incorrectChars, string(guess[pos]))
			} else if char == "P" {
				correctCharsWrongPositions[string(guess[pos])] = append(correctCharsWrongPositions[string(guess[pos])], pos)
			}
		}

		// Only required for debugging
		fmt.Println("Correct Letters:", correct)
		fmt.Println("Incorrect Letters:", incorrectChars)
		fmt.Println("Correct Letters But Wrong Positions:", correctCharsWrongPositions)

		var candidates []string
		for _, word := range possibleWords {

			if !allCorrectCharsInWord(correct, word) {
				continue
			}

			if wordContainsIncorrectChar(word, incorrectChars) {
				continue
			}

			// Check if it contains the characters in wrong position, skip if it does not
			// check if string "w" contains correct chars but wrong position
			doesNotContainChar := false
			containsCorrectCharInWrongPos := false
			for correctChar, wrongPositions := range correctCharsWrongPositions {
				if !wordContainsChar(word, correctChar) {
					doesNotContainChar = true
					break
				} else {

					// Contains the correct character, check if
					// correct character is in the wrong position
					containsCorrectCharInWrongPos = false
					for _, wrongPos := range wrongPositions {
						if string(word[wrongPos]) == correctChar {
							containsCorrectCharInWrongPos = true
							break
						}
					}
					if containsCorrectCharInWrongPos {
						break
					}
				}
			}
			if doesNotContainChar {
				continue
			}
			if containsCorrectCharInWrongPos {
				continue
			}

			candidates = append(candidates, word)
		}

		// Reduce possibleWords to candidates
		possibleWords = make([]string, len(candidates))
		copy(possibleWords, candidates)
		fmt.Println(len(possibleWords), "Possible Words:", possibleWords)

		// No result or only 1 result
		if len(possibleWords) <= 1 {
			break
		}
	}
}

func allCorrectCharsInWord(correctChars [5]string, word string) bool {
	// All correct chars match word
	for pos, c := range correctChars {
		char := string(c)
		if char != "" && char != string(word[pos]) {
			return false
		}
	}

	return true
}

func wordContainsChar(str string, char string) bool {
	for _, c := range str {
		if string(c) == char {
			return true
		}
	}
	return false
}

func wordContainsIncorrectChar(word string, incorrectChars []string) bool {
	// check if string "w" contains a character that IS NOT IN WORD
	for _, c := range incorrectChars {
		if wordContainsChar(word, c) {
			return true
		}
	}
	return false
}
