package main

import (
	"fmt"
	"testing"
)

func TestAdventocode(t *testing.T) {
	t.Run("testing if the string contains three vowels", func(t *testing.T) {
		myString := "ugknbfddgicrmopn"
		got := CheckIfContainsThreeVowels(myString)
		want := 1

		if got != want {
			t.Errorf("The string '%q' don't contains a minimum quantity vowels\n", myString)
		}
	})
	t.Run("verify if contains double letters", func(t *testing.T) {
		myString := "ugknbfddgicrmopn"
		got := CheckIfContainsDoubleLetter(myString)
		want := 1

		if got != want {
			t.Errorf("The string '%q' fon't contains double letter\n", myString)
		}
	})
	t.Run("verify some proibide sequences", func(t *testing.T) {
		myString := "ugknbfddgicrmopn"
		got := CheckIfDontContainsACertalySequence(myString)
		want := 1
		if want != got {
			t.Errorf("The string '%q' contains a certaly serquence that is not permited", myString)
		}
	})
	t.Run("Verify NiceString in slice of the strings", func(t *testing.T) {
		mySlice := []string{"ugknbfddgicrmopn", "rthkunfaakmwmush", "lbxlcixxcztddlam", "lbxlcixxycztddlam"}
		got := VerifyNiceString(mySlice)
		want := 2
		if len(got) != want {
			t.Errorf("GOT %d  WANT %d", len(got), want)
		}
	})
}

func TestNewRulerNiceString(t *testing.T) {
	t.Run("Check if have two pair equals", func(t *testing.T) {
		myString := "aabcdefgaa"
		got := CheckSameTwoPair(myString)
		want := 1
		if got != want {
			t.Errorf("This string '%q' don't containt two pair equals", myString)
		}
	})

	t.Run("check if have two same letter and anoother letter between them", func(t *testing.T) {
		myString := "xxx"
		got := CheckOneLetterBetweenAnotherBothLetter(myString)
		want := 1
		if got != want {
			t.Errorf("This string '%q' havent a letter between two both another letter like 'xyx' ", myString)
		}
	})
	t.Run("Verify a new NiceString", func(t *testing.T) {
		mySlice := []string{"aeiouaeiouaeieu", "qjhvhtzxzqqjkmpb", "xxyxx", "uurcxstgmygtbstg", "ieodomkazucvgmuy"}
		got := VerifyNiceString2(mySlice)
		want := 3
		fmt.Println(got)
		if len(got) != want {
			t.Errorf("GOT %d  WANT %d", len(got), want)
		}

	})
}
