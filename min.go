package main

import (
	"fmt"
	"strings"
)

const (
	letters       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterCodeMap = "UMRSQPBOLEXTZYAKJVCNHDWGIF"
)

func substituteCipher(input string) string {
	output := ""
	for _, char := range input {
		if 'A' <= char && char <= 'Z' {
			index := strings.IndexRune(letters, char)
			output += string(letterCodeMap[index])
		} else {
			output += string(char)
		}
	}
	return output
}

func main() {
	input := "HELLO"
	encrypted := substituteCipher(input)
	fmt.Println("Encrypted:", encrypted)
}
