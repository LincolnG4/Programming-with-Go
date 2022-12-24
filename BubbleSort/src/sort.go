package main

import (
	"fmt"
)

var userInput int
var slice []int

func Swap(slice []int, index int) {
	temp := slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = temp
}

func BubbleSort(slice []int) {

	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}

	}
}
func UserInputArray(userInput int) []int {
	slice := make([]int, 10)

	for i := 0; i <= 9; i++ {
		fmt.Printf("Insert #%d number:", i+1)
		fmt.Scan(&userInput)
		slice[i] = userInput
	}

	return slice
}

func main() {
	slice = UserInputArray(userInput)

	BubbleSort(slice)
	fmt.Println(slice)

}
