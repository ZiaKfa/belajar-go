package main

import "fmt"

func main() {
	//constant variable must be initialized
	const phi = 3.14
	fmt.Println(phi, "is a constant")
	//phi = 3.14159265358979323846264338327950288419716939937510582097494459
	const temp = 69
	//constant variable can be initialized without using it
	const (
		firstName = "John"
		lastName  = "Doe"
	)
}
