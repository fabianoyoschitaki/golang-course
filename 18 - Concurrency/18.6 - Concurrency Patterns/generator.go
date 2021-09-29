package main

import (
	"fmt"
	"time"
)

func main() {
	// we don't need to worry about make, creating channels etc
	channel := write("Hello World")
	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

// this is a generator pattern, encapsulates a goroutine and returns a communication channel
func write(text string) <-chan string {
	// creating channel
	channel := make(chan string)
	go func() {
		for {
			// adding text to channel
			channel <- fmt.Sprintf("Received value: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return channel
}
