package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

var readPerson Person
var filePath string
var file *os.File
var err error
var personSlice []Person

func main() {
	fmt.Print("Input Txt File Path:")
	fmt.Scan(&filePath)

	file, err = os.Open(filePath)
	if err != nil {
		fmt.Println("unable to read input file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	i := 0
	for scanner.Scan() {
		readPerson.fname = strings.Split(scanner.Text(), " ")[0]
		readPerson.lname = strings.Split(scanner.Text(), " ")[1]
		personSlice = append(personSlice, readPerson)
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	for _, value := range personSlice {
		if value.fname == "" && value.lname == "" {
			break
		}
		fmt.Println(value.fname, value.lname)
	}
}
