package main

import (
	"fmt"
	"time"
)

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func main() {
	go write("Hello 1") // goroutine: execute this function, but don't wait for it to continue
	write("Hello 2")    // if you put go here, nothing will happen because next line the program is over, so program will close
	// program finishes here
}
