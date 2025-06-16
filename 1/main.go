package main

import (
	"fmt"
	"os"
)

const (
	LEFT_PARENTHESIS  = 40
	RIGHT_PARENTHESIS = 41
)

func ParenthesisCountFloor(list []uint8) (int, error) {
	result := 0
	for i := range list {
		if list[i] == LEFT_PARENTHESIS {
			result++
		}
		if list[i] == RIGHT_PARENTHESIS {
			result--
		}
	}
	return result, nil
}

func CheckWhenBackTheFistFloor(list []uint8) int {
	floor := 0
	for i := range list {
		if list[i] == LEFT_PARENTHESIS {
			floor++
		}
		if list[i] == RIGHT_PARENTHESIS {
			floor--
		}
		if floor == -1 {
			return i + 1
		}
	}
	return 0
}

func main() {
	dataFile, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "a error has ocurred, %v\n", err)
	}

	parenthesisList := dataFile
	countFloor, err := ParenthesisCountFloor(parenthesisList)
	backtoFirstFloot := CheckWhenBackTheFistFloor(parenthesisList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "a error has ocurred, %v\n", err)
	}
	fmt.Printf("Last Floor: %d\n",countFloor)
	fmt.Printf("When back to the First Floor: %d\n",backtoFirstFloot)
}
