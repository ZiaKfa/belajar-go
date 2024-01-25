package main

import "fmt"

func main() {
	//if statement
	username := "admin"
	password := "admin"
	if username == "admin" && password == "admin" {
		fmt.Println("You are logged in")
	} else if username == "admin" && password != "admin" {
		fmt.Println("Wrong password")
	} else if username != "admin" && password == "admin" {
		fmt.Println("Wrong username")
	} else {
		fmt.Println("Wrong username and password")
	}

	//short statement
	word := "Hello"
	// if len(word) < 5 {
	// 	fmt.Println("Word is too short")
	// } else {
	// 	fmt.Println("Word is good")
	// }
	//the code above is the same as this
	if length := len(word); length < 5 {
		fmt.Println("Word is too short")
	} else {
		fmt.Println("Word is good")
	}

	//switch statement
	dayNum := 2
	switch dayNum {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("The day is not monday, tuesday, or wednesday")
	}

	//switch statement with short statement
	switch name := "John"; name {
	case "John":
		fmt.Println("Hello John")
	case "Jane":
		fmt.Println("Hello Jane")
	default:
		fmt.Println("Who are you?")
	}

	//switch statement with condition (like if-else)
	switch word := "Hello"; len(word) < 5 {
	case true:
		fmt.Println("Word is too short")
	case false:
		fmt.Println("Word is good")
	}

	//switch statement without condition (like if-else)
	word = "Hello!"
	length := len(word)
	switch {
	case length < 5:
		fmt.Println("Word is too short")
	case length > 5:
		fmt.Println("Word is good")
	default:
		fmt.Println("Word is exactly 5 characters")
	}
}
