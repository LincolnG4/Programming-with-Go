package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputString string

	fmt.Print("Enter a string: ")
	fmt.Scan(&inputString)

	inputString = strings.ToLower(inputString)
	if strings.HasPrefix(inputString, "i") && strings.Contains(inputString, "a") && strings.HasSuffix(inputString, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
