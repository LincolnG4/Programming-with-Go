package main

import (
	"fmt"
	"time"
)

/* A race condition occurs when multiple threads try to access and modify the same data (memory address).
E.g., if one thread tries to increase an integer and another thread tries to read it,
this will cause a race condition. On the other hand, there won't be a race condition,
if the variable is read-only. In golang, threads are created implicitly when Goroutines are used.*/

var sharedVar int = 0

func varIncrease() {
	/* A goroutine reads the variable named sharedVar
	 */
	for i := 1; i < 50; i++ {
		fmt.Print("Channel 1 ->")
		fmt.Println("Shared variable= ", sharedVar)
		sharedVar++
	}

}

func varReader() {
	// goroutine writes to the same variable by increasing its value.

	for i := 1; i < 50; i++ {
		fmt.Print("Channel 2 ->")
		fmt.Println("Shared variable= ", sharedVar)
	}
}

func main() {
	/* We are not sure as to which go routine
	will execute first and which one second as they are different threads.
	Since they are sharing the same variable 'sharedVar', we need to make sure that they are different threads
	what is impossible to do cause it is not deterministic.
	with you run this code multiple times, sometimes final results channel 2 will be 49, sometimes 0
	*/

	go varIncrease()
	go varReader()
	time.Sleep(time.Second)
}
