package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	userMap := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Print a name: ")
	if scanner.Scan() {
		userInput := scanner.Text()
		userMap["name"] = userInput
	}

	fmt.Print("Print a addr: ")
	if scanner.Scan() {
		userInput := scanner.Text()
		userMap["address"] = userInput

	}

	barr, _ := json.Marshal(userMap)

	fmt.Printf("String version: %s", string(barr))

}
