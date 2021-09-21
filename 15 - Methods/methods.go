package main

import "fmt"

// we can create actions inside user.
// go is not a OOP language, but it supports methods
type user struct {
	name string
	age  uint8
}

// not a func! it's a method. Attached to user
func (u user) toString() {
	fmt.Printf("User %s with age %d\n", u.name, u.age)
}

func (u user) isOlderThan(age uint8) bool {
	return u.age > age
}

// changing value of object
// IMPORTANT does not need dereferencing
func (u *user) makeAnniversary() {
	u.age++
}

func main() {
	user1 := user{"Fabiano", 10}
	user1.toString()
	fmt.Println(user1.isOlderThan(9))

	user2 := user{name: "Joao"}
	user2.toString()
	user2.makeAnniversary()
	user2.toString()
}
