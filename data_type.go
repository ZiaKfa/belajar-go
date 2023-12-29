package main

import "fmt"

func main() {
	//String
	var helo string = "Hello World"
	fmt.Println(helo + " is a string")
	fmt.Println(helo[0], "is a binary byte of first index char from ", helo, " string")
	fmt.Println(len(helo), "is the length of string ", helo)

	//Number
	var num int32 = 10
	fmt.Println(num, "is an integer")
	var num2 int64 = 1000000000000000000
	fmt.Println(num2, "is a 64 bit integer ")

	//Float
	var phiShort float32 = 3.14
	fmt.Println(phiShort, "is a 32 bit float ")
	var phiLong float64 = 3.14159265358979323846264338327950288419716939937510582097494459
	fmt.Println(phiLong, "is a 64 bit float ")

	//Boolean
	var betul bool = true
	fmt.Println(betul, "is a boolean")
	var salah bool = false
	fmt.Println(salah, "is a boolean")
}
