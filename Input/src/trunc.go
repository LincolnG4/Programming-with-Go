package main

import (
	"fmt"
)

func main() {
	var inputUser float64
	fmt.Print("Input a float number: ")
	fmt.Scan(&inputUser)
	var truncate int = int(inputUser)
	fmt.Println(truncate)
}
