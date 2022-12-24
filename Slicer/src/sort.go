package main

import (
	"fmt"
	"sort"
	"strconv"
)

var userInput string = ""
var err string = ""

func main() {
	slice := make([]int, 3)
	for userInput != "x" {
		fmt.Println("Enter an interger. Enter 'x' to exit...")
		fmt.Scan(&userInput)
		if userInput != "x" {
			intUserInput, err := strconv.Atoi(userInput)
			if err == nil {
				slice = append(slice, intUserInput)
				sort.Ints(slice)
				fmt.Println(slice)
			} else {
				fmt.Println("Invalid input")
			}
		} else {
			break
		}

	}

}
