package main

import "fmt"

//function without parameter and return value
func sayHello() {
	fmt.Println("Hello World")
}

//function with parameter and without return value
func printName(name string) {
	fmt.Println("Hello", name)
}

//function with parameter and return value
func getName(name string) string /*<== this is the return type*/ {
	return "Hello " + name
}

//function with multiple return value
//unlike other programming language, Go function can return multiple value
func getFullName(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

func getFullNameArray(firstName string, lastName string) [2]string {
	return [2]string{firstName, lastName}
}

func main() {
	sayHello()
	printName("John Doe")
	fmt.Println(getName("John Doe"))
	/* the return variable here can be assigned to ( _ ) if the value is not used */
	firstName, lastName := getFullName("John", "Doe")
	fmt.Println(firstName, lastName)
	fullName := getFullNameArray("John", "Doe")
	fmt.Println(fullName[0])
	fmt.Println(fullName[1])
}
