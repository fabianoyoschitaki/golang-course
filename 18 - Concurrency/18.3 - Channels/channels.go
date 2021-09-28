package main

import (
	"fmt"
	"time"
)

// channel : communication channels most common used to sync goroutines in a better way
func main() {
	channel := make(chan string) // only string types (send and retrieve)
	go write("Hello 1", channel) // 1. 'go' to make it a goroutine

	fmt.Println("After go write call")

	for { // this will cause a deadlock because we're waiting for a message to the channel which will never be sent: fatal error: all goroutines are asleep - deadlock!
		msg, isOpen := <-channel // 2. receiving the text in the channel. wait for it! otherwise nothing below runs
		fmt.Println(msg)         // 4. printing message sent to channel
		if !isOpen {
			break
		}
	}

	// a smarter way to do the same stuff as above
	// for msg := range channel {
	// 	fmt.Println(msg)
	// }
	fmt.Println("End of execution")
}

func write(text string, channel chan string) { // channel is used to send and retrieve data
	time.Sleep(time.Second * 2)
	for i := 0; i < 5; i++ {
		channel <- text // 3. sending a value (string in this case) to the channel
		time.Sleep(time.Second)
	}
	// this tells we're not sending messages to the channel anymore
	close(channel)
}

// a channel can send and receive data. they block the program execution
