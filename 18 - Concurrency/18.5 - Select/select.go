package main

import (
	"fmt"
	"time"
)

func main() {
	// creating 2 goroutines which will send to different channels
	channel1, channel2 := make(chan string), make(chan string)

	// anon function
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "Canal 1"
		}
	}()

	// anon function
	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Canal 2"
		}
	}()

	// endless for
	for {
		// unnecessary delay, because it waits for channel2. how to solve?
		// messageFromChannel1 := <-channel1
		// fmt.Println(messageFromChannel1)

		// messageFromChannel2 := <-channel2
		// fmt.Println(messageFromChannel2)

		select { // if there are messages from channel1 or channel2, do it
		case messageFromChannel1 := <-channel1:
			fmt.Println(messageFromChannel1)
		case messageFromChannel2 := <-channel2:
			fmt.Println(messageFromChannel2)
		}
	}
}
