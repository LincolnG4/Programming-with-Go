package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

var userAnimalInput string
var userMethodInput string

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	animalMap := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	for {
		fmt.Print("> ")

		fmt.Scan(&userAnimalInput)
		fmt.Scan(&userMethodInput)

		userAnimalInput = strings.ToLower(userAnimalInput)
		userMethodInput = strings.ToLower(userMethodInput)

		if strings.ToLower(userMethodInput) == "eat" {
			animalMap[userAnimalInput].Eat()
		} else if strings.ToLower(userMethodInput) == "move" {
			animalMap[userAnimalInput].Move()
		} else if strings.ToLower(userMethodInput) == "speak" {
			animalMap[userAnimalInput].Speak()
		} else {
			fmt.Printf("Animal %s or method %s is not listed \n---> Listed animal: Cow, Bird and Snake\n---> Listed methods: Eat, Move, Speak\n", userAnimalInput, userMethodInput)
		}
	}
}
