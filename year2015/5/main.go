package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	mySliceString := getStringsByFile("./input.txt")
	niceStrings := VerifyNiceString(mySliceString)
	niceStrings2 :=VerifyNiceString2(mySliceString) 
	fmt.Printf("You have %d nice strings\n", len(niceStrings))
	fmt.Printf("You have %d nice strings 2.0\n", len(niceStrings2))
}

func getStringsByFile(filePath string) (sliceString []string) {

	fileContent, _ := os.ReadFile(filePath)
	for line := range strings.SplitSeq(string(fileContent), "\n") {
		sliceString = append(sliceString, line)
	}
	return
}

func CheckIfContainsThreeVowels(s string) int {
	checherVowels := 0

	if len(s) < 3 {
		return 0
	}
	for word := range s {
		if string(s[word]) == "a" || string(s[word]) == "e" || string(s[word]) == "i" || string(s[word]) == "o" || string(s[word]) == "u" {
			checherVowels++
		}
		if checherVowels >= 3 {
			return 1
		}
	}
	return 0
}

func CheckIfContainsDoubleLetter(s string) int {

	prevLetter := ""
	for i, letter := range s {
		if i-1 == -1 {
			prevLetter = string(letter)
		} else {
			if prevLetter == string(letter) {
				return 1
			}
			prevLetter = string(letter)
		}
	}
	return 0
}

func CheckIfDontContainsACertalySequence(s string) int {
	sequenceNotPermited := []string{"ab", "cd", "pq", "xy"}
	prevLetter := ""
	for i, letter := range s {
		if i-1 == -1 {
			prevLetter = string(letter)
		} else {
			if slices.Contains(sequenceNotPermited, prevLetter+string(letter)) {
				return 0
			}
			prevLetter = string(letter)
		}
	}
	return 1
}

func VerifyNiceString(myStrings []string) (myNiceString []string) {
	tempNumber := 0
	for _, s := range myStrings {
		tempNumber = 0
		tempNumber += CheckIfContainsThreeVowels(s)
		tempNumber += CheckIfContainsDoubleLetter(s)
		tempNumber += CheckIfDontContainsACertalySequence(s)

		if tempNumber == 3 {
			myNiceString = append(myNiceString, s)
		}
	}
	return
}

func CheckSameTwoPair(s string) int {
	prevLetter := ""
	for i, letter := range s {
		if i-1 == -1 {
			prevLetter = string(letter)
		} else {
			numberOfPair := strings.Count(s, prevLetter+string(letter))
			if numberOfPair > 1 {
				return 1
			}
			prevLetter = string(letter)
		}
	}
	return 0
}

func CheckOneLetterBetweenAnotherBothLetter(s string) int {
	prevLetter := ""
	for i, letter := range s {
		if i-1 == -1 {
			prevLetter = string(letter)
		} else {
			if strings.Contains(s, prevLetter+string(letter)+prevLetter) {
				return 1
			}
			prevLetter = string(letter)
		}
	}
	return 0
}

func VerifyNiceString2(sliceString []string) (myNiceString []string) {
	tempNumber := 0
	for _, s := range sliceString {
		tempNumber = 0
		tempNumber += CheckSameTwoPair(s)
		tempNumber += CheckOneLetterBetweenAnotherBothLetter(s)

		if tempNumber == 2 {
			myNiceString = append(myNiceString, s)
		}
	}
	return
}
