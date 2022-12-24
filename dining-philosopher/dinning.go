package main

import (
	"fmt"
	"sync"
	"time"
)

type CS struct{ sync.Mutex }

type Ph struct {
	id       int
	leftCS   *CS
	rightCS  *CS
	eatTimes int
	allowed  chan bool
}

type host struct {
	requestChannel    chan int
	abortChannel      chan bool
	finishedChannel   chan int
	onlineUsers       map[int]bool
	numberOnlineUsers int
}

func (h *host) requestEat(Philosophers []*Ph) {
	for {

		select {
		case id := <-h.requestChannel:
			if h.numberOnlineUsers < 2 || h.onlineUsers[id] == true {
				fmt.Println("(HOST) ID ", Philosophers[id-1].id, " --> Eat request: ALLOWED")
				h.numberOnlineUsers++
				h.onlineUsers[id] = true
				Philosophers[id-1].allowed <- true
			} else {
				fmt.Println("(HOST) ID ", Philosophers[id-1].id, " --> Eat request: DENIED")
				Philosophers[id-1].allowed <- false
			}

		case id := <-h.finishedChannel:
			fmt.Println("(HOST) ID ", Philosophers[id-1].id, " --> Finished eat ")
			h.numberOnlineUsers--
			h.onlineUsers[id] = false

		case <-h.abortChannel:
			fmt.Println("(HOST) Closing channel")
			return
		}
	}

}

func (p *Ph) eat(wg *sync.WaitGroup, h *host) {

	for p.eatTimes != 3 {
		fmt.Println("GOROUTINE.ID:", p.id, " --> Requesting permission")

		// Request Permissions
		h.requestChannel <- p.id
		allowedFlag := <-p.allowed
		if allowedFlag == true {
			fmt.Println("GOROUTINE.ID:", p.id, " --> Queue allowed: ", h.onlineUsers)
			p.leftCS.Lock()
			p.rightCS.Lock()

			fmt.Println("GOROUTINE.ID:", p.id, " --> Starting to eat ")

			time.Sleep(time.Second)

			p.leftCS.Unlock()
			p.rightCS.Unlock()

			fmt.Println("GOROUTINE.ID:", p.id, " --> Finished eat  ")
			h.finishedChannel <- p.id

			p.eatTimes++
		} else {
			time.Sleep(time.Millisecond * 500)
		}

	}

	fmt.Println("GOROUTINE.ID:", p.id, " DONE --> Times: ", p.eatTimes)
	wg.Done()
}

var wg sync.WaitGroup
var once sync.Once

func main() {

	usersmap := map[int]bool{
		1: false,
		2: false,
		3: false,
		4: false,
		5: false,
	}
	// Create host
	hostRequest := &host{make(chan int), make(chan bool), make(chan int), usersmap, 0}

	// Create Sticks list
	Sticks := make([]*CS, 5)

	for i := 0; i < 5; i++ {
		Sticks[i] = new(CS)
	}

	// Create Philosophers list
	Philosophers := make([]*Ph, 5)

	for i := 0; i < 5; i++ {
		Philosophers[i] = &Ph{i + 1, Sticks[i], Sticks[(i+1)%5], 0, make(chan bool)}
	}

	// once
	go hostRequest.requestEat(Philosophers)

	// Start Routine
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go Philosophers[i].eat(&wg, hostRequest)
	}

	wg.Wait()

	hostRequest.abortChannel <- true
}

func idInSlice(checkid int, list []int) bool {
	for _, id := range list {
		if id == checkid {
			return true
		}
	}
	return false
}
