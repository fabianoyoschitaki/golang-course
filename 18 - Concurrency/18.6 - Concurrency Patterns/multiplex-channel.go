package main

import (
	"fmt"
	"math/rand"
	"time"
)

// multiplex: join 2 or more channels
func main() {
	outputChannel := multiplex(write("This is from channel 1"), write("This is from channel 2"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-outputChannel)
	}
}

func multiplex(channel1 <-chan string, channel2 <-chan string) <-chan string {
	outputChannel := make(chan string)

	go func() {
		for {
			select {
			case msg := <-channel1:
				outputChannel <- msg
			case msg := <-channel2:
				outputChannel <- msg
			}
		}
	}()
	return outputChannel
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Received text: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
