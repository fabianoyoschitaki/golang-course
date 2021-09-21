package main

import "fmt"

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("Successfully recovered!")
	}
}

func isStudentApproved(n1, n2 float32) bool {
	defer recoverExecution()
	defer fmt.Println("Average is calculated.")

	fmt.Println("Checking if student is approved")
	average := (n1 + n2) / 2
	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}
	// call all deferred functions and stop execution
	panic("Average is 6! Stop everything")
}

// 1. run isStudentApproved
// 2. Checking if student is approved
// 3. Average is calculated.
// 4. Successfully recovered!
// 5. false
// 6. After execution
func main() {
	fmt.Println(isStudentApproved(6, 6)) // if 6,6 panic will be run and line below will not run
	fmt.Println("After execution")
}
