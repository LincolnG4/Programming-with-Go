package main

import (
	"fmt"
	"reflect"
)

var x int = 1
var pointer *int
var y int

func main() {
	pointer = &x
	y = *pointer
	fmt.Println(x, pointer, y)
	x = 2
	y = *pointer
	fmt.Println(x, pointer, y)
	p := new(int)
	*p = 3
	fmt.Printf("p: %s\n", reflect.TypeOf(p))
	fmt.Printf("pointer: %s\n", reflect.TypeOf(pointer))
	fmt.Printf("x: %s\n", reflect.TypeOf(x))

}
