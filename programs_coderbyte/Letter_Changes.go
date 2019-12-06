package main

import (
	"fmt"
	"strings"
)

var vowels = `aeiou`
var alphabet = `abcdefghiklmnopqrstvxyz`

func LetterChange(str string) string {

	var newString = ``
	for _, char := range str {
		currentChar := string(char)
		charCheck := strings.ToLower(currentChar)

		if charCheck == `z` {
			currentChar = `a`
		} else {
			currentChar = string(int(char) + 1)
		}

		newString += currentChar

	}
	return newString
}

func main() {
	var tring string
	fmt.Scan(&tring)

	newstr := LetterChange(tring)
	fmt.Println(newstr)
}
