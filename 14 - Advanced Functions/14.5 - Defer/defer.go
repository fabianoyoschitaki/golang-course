package main

import "fmt"

func func1() {
	fmt.Println("This is func1")
}

func func2() {
	fmt.Println("This is func2")
}

func isStudentApproved(n1, n2 float32) bool {
	defer fmt.Println("Average is calculated.") // to avoid duplicates, we can put defer to be executed immediately before the end!
	fmt.Println("Checking if student is approved")
	average := (n1 + n2) / 2
	if average >= 6 {
		// fmt.Println("Average is calculated.")
		return true
	}
	// fmt.Println("Average is calculated.")
	return false
}

// defer is very useful for DBs, you execute statements and in the end you close the connection with defer
func main() {
	defer func1() // this will make defer this function execution to the latest moment, in this case the end of main function
	func2()
	fmt.Println(isStudentApproved(7, 8)) // even having return
}
