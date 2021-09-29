package main

import "fmt"

func main() {
	// opening channels
	tasks := make(chan int, 45)
	results := make(chan int, 45)

	// the more the better. work pools concurrency pattern
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	// sending data to tasks
	for i := 0; i < 45; i++ {
		tasks <- i
	}
	close(tasks)

	// reading data from results
	for i := 0; i < 45; i++ {
		result := <-results
		fmt.Println(result)
	}

}

// making tasks channel only to receive, while results send data
func worker(tasks <-chan int, results chan<- int) {
	for number := range tasks {
		results <- fibonacci(number)
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
