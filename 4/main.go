package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func FindAdventCoin(secretKey string) (int, string) {
	return findAdventCoin(secretKey, "00000")
}
func findAdventCoin(secretKey, prefix string) (int, string) {
	counter := 0
	found := false

	for !found {
		input := fmt.Sprint(secretKey, counter)
		hash := md5.Sum([]byte(input))
		hashHex := fmt.Sprintf("%x", hash)

		if strings.HasPrefix(hashHex[:len(prefix)], prefix) {
			return counter, string(hash[:])
		} else {
			counter++
		}
	}

	return 0, ""
}

func main() {
	secretKey := "iwrupvqb"
	result, coin := FindAdventCoin(secretKey)
	fmt.Printf("Your secret Key: %q\nYour answer: %d\nYour AdventCoin: %x\n", secretKey, result, coin)
	fmt.Println()
	result, coin = findAdventCoin(secretKey, "000000")
	fmt.Printf("Your secret Key: %q\nYour answer: %d\nYour AdventCoin: %x\n", secretKey, result, coin)
}
