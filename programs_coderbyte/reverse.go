package main

import (
	"fmt"
)

func ReverseString(str string) string {
	var result = ``
	for i := (len(str) - 1); i >= 0; i-- {
		result += string(str[i])

	}
	return result
}

// strings.Index()

func main() {
	var data string
	fmt.Scan(&data)
	fmt.Println(ReverseString(data))
}
