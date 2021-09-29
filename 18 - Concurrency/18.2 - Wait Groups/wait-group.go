package main

import (
	"fmt"
	"sync"
	"time"
)

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

// WaitGroup is a way to synchronize goroutines: to wait them to finish

func main() {
	var waitGroup sync.WaitGroup // creating waitgroup
	waitGroup.Add(2)             // 2 goroutines. next 2 ones

	// anonymous function - go routine 1
	go func() {
		write("Hello 1")
		waitGroup.Done() // this will remove 1 from the counter (of 2 goroutines)
	}()
	// anonymous function - go routine 2
	go func() {
		write("Hello 2")
		waitGroup.Done() // this will remove 1 from the counter (of 2 goroutines)
	}()

	waitGroup.Wait() // does not continue until 2 goroutines finish
	fmt.Println("Now we can continue")

}
