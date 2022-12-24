package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
)

var userArray []int
var userInput string
var i int = 0
var wg sync.WaitGroup

func sortSlice(splitSlice []int, c chan []int, id int) {
	sort.Ints(splitSlice)
	fmt.Println("GoRoutine ", id, ": ", splitSlice)
	c <- splitSlice
	wg.Done()

}

func main() {
	start := time.Now()
	// Code to measure
	duration := time.Since(start)
	c := make(chan []int, 4)

	for {
		fmt.Println("Enter an interger. Enter 'Q' to exit...")

		fmt.Scan(&userInput)
		if userInput != "Q" {
			inputString, _ := strconv.Atoi(userInput)
			userArray = append(userArray, inputString)
			i++
		} else {
			break
		}
	}

	avg := float64(len(userArray)) / 4.0
	last := 0.0
	id := 1

	wg.Add(4)

	for last < float64(len(userArray)) {
		go sortSlice(userArray[int(last):int(last+avg)], c, id)
		last += avg
		id++
	}
	wg.Wait()

	slice1, slice2, slice3, slice4 := newFunction(c)

	slicesPackage := [][]int{slice1, slice2, slice3, slice4}
	concatSlice := []int{}
	for _, slice := range slicesPackage {
		concatSlice = append(concatSlice, slice...)
	}
	sort.Ints(concatSlice)

	fmt.Println(concatSlice)
	fmt.Println(duration)

}

func newFunction(c chan []int) ([]int, []int, []int, []int) {
	slice1 := <-c
	slice2 := <-c
	slice3 := <-c
	slice4 := <-c
	return slice1, slice2, slice3, slice4
}
