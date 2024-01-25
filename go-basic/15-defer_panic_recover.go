package main

//unlike other programming language, Go doesn't have try-catch block
//but Go has defer, panic, and recover function to handle error

import "fmt"

func endApp() {
	fmt.Println("app stopped")
	errMessage := recover() //recover() function will catch the panic() function
	//make sure to use recover() function inside defer function, otherwise it will not work
	fmt.Println("error message :", errMessage)
}

func loadApp(error bool) {
	defer endApp() //this function will be executed after runApp() function finished, even if there is an error
	if error {
		panic("Oops error") //this will stop the program and execute defer function
	}
	fmt.Println("app loaded")
}

func main() {
	loadApp(true)
	fmt.Println("app is running")
}
