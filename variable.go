package main

import "fmt"

func main() {
	//complete variable declaration
	var hello string = "Hello World"
	fmt.Println(hello + " is a string")

	//variable declaration without type inference, automatically inferred
	var name = "John Doe"
	fmt.Println("Hello", name)
	name = "Jane Doe"
	fmt.Println("Hello", name)

	//variable declaration without var keyword, automatically inferred
	age := 20
	fmt.Println("Age", age)
	age = 30
	fmt.Println("Age", age)

	//multiple variable declaration
	var (
		firstName = "John"
		lastName  = "Doe"
	)
	fmt.Println("Hello", firstName, lastName)
}
