package main

import (
	"fmt"
	"math"
)

var accelerationInput float64
var velocityInput float64
var displacementInput float64
var t float64
var displacementResult float64

func GenDisplaceFn(accelerationInput float64, velocityInput float64, displacementInput float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*accelerationInput*math.Pow(t, 2) + velocityInput*t + displacementInput
	}
	return fn
}

func main() {

	fmt.Print("Insert acceleration: ")
	fmt.Scan(&accelerationInput)

	fmt.Print("Insert initial velocity: ")
	fmt.Scan(&velocityInput)

	fmt.Print("Insert initial displacement: ")
	fmt.Scan(&displacementInput)

	displacementFunction := GenDisplaceFn(accelerationInput, velocityInput, displacementInput)

	fmt.Print("Insert time: ")
	fmt.Scan(&t)

	// s = ½ a t2 + vot + so
	fmt.Println(displacementFunction(t))
}
