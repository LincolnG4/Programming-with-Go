package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
}

func (d Cow) Eat() {
	fmt.Println("grass")
}

func (d Cow) Move() {
	fmt.Println("walk")
}

func (d Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
}

func (d Bird) Eat() {
	fmt.Println("worms")
}

func (d Bird) Move() {
	fmt.Println("fly")
}

func (d Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
}

func (d Snake) Eat() {
	fmt.Println("mice")
}

func (d Snake) Move() {
	fmt.Println("slither")
}

func (d Snake) Speak() {
	fmt.Println("hsss")
}

var userMenuInput string
var userAnimalInput string
var userMethodInput string
var userAnimalTypeInput string

func main() {
	animalsTable := make(map[string]Animal)

	for {
		fmt.Println("Type 'newanimal' or 'query'")
		fmt.Print("> ")

		fmt.Scan(&userMenuInput)

		if userMenuInput == "newanimal" {

			userNewAnimal := "Unknown"
			fmt.Println("Creating new animal (default: 'Unknown')")
			fmt.Println("Animal name:")
			fmt.Print("> ")
			fmt.Scan(&userNewAnimal)

			fmt.Println("Animal type (cow, bird or snake):")
			fmt.Print("> ")
			fmt.Scan(&userAnimalTypeInput)

			switch strings.ToLower(userAnimalTypeInput) {
			case "cow":
				animalsTable[userNewAnimal] = new(Cow)
				fmt.Println("Created it!")
			case "bird":
				animalsTable[userNewAnimal] = new(Bird)
				fmt.Println("Created it!")
			case "snake":
				animalsTable[userNewAnimal] = new(Snake)
				fmt.Println("Created it!")
			default:
				fmt.Println("Type doens't exist")
			}

		} else {
			fmt.Println("Query animal")
			fmt.Print("> ")
			fmt.Scan(&userAnimalInput)
			fmt.Scan(&userMethodInput)

			userAnimalInput = strings.ToLower(userAnimalInput)
			userMethodInput = strings.ToLower(userMethodInput)

			if strings.ToLower(userMethodInput) == "eat" {
				animalsTable[userAnimalInput].Eat()
			} else if strings.ToLower(userMethodInput) == "move" {
				animalsTable[userAnimalInput].Move()
			} else if strings.ToLower(userMethodInput) == "speak" {
				animalsTable[userAnimalInput].Speak()
			} else {
				fmt.Printf("Animal %s or method %s is not listed \n---> Listed animal: Cow, Bird and Snake\n---> Listed methods: Eat, Move, Speak\n", userAnimalInput, userMethodInput)
			}
		}

	}
}
