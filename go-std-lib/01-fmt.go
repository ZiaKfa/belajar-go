package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	firstName := "John"
	lastName := "Doe"
	fmt.Println("Hello '", firstName, lastName, "'") //will have additional spaces

	//using string formatting, similar to C
	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
