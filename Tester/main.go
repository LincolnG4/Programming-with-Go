// Developed by: Leandro Zaneratto
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var item = make(chan bool, 2)

type Shopstick struct{ sync.Mutex }

type Philosopher struct {
	identifier     int
	leftShopstick  *Shopstick
	rightShopstick *Shopstick
}

func (philosopher Philosopher) eat() {

	for i := 0; i < 3; i++ {
		item <- true
		philosopher.leftShopstick.Lock()
		philosopher.rightShopstick.Lock()
		fmt.Println("Starting to eat Philosopher: ", philosopher.identifier+1)
		time.Sleep(2 * time.Second)
		fmt.Println("finishing eating Philosopher: ", philosopher.identifier+1)
		philosopher.leftShopstick.Unlock()
		philosopher.rightShopstick.Unlock()
		<-item
	}
	wg.Done()
}

func main() {

	shopsticks := make([]*Shopstick, 5)
	for i := 0; i < 5; i++ {
		shopsticks[i] = new(Shopstick)
	}
	Philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		Philosophers[i] = &Philosopher{i, shopsticks[i], shopsticks[(i+1)%5]}
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Philosophers[i].eat()
	}
	wg.Wait()
}
