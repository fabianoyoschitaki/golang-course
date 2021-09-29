package main

import "fmt"

func main() {
	channel := make(chan string, 2) // 2 is buffer. this means it will only get blocked when reaching max capacity
	channel <- "Hello world"        // gets blocked here. that's why we use channels in different functions
	channel <- "Hello world 2"      // gets blocked here. that's why we use channels in different functions
	// channel <- "Hello world 3"      // deadlock, capacity is 2

	msg := <-channel
	msg2 := <-channel
	// msg3 := <-channel // capacity is 2: fatal error: all goroutines are asleep - deadlock!

	fmt.Println(msg)
	fmt.Println(msg2)
	// fmt.Println(msg3)
}
