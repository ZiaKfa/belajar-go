package main

import "fmt"

func main() {
	var helo string = "Hello World"
	fmt.Println(helo + " adalah string")
	fmt.Println(helo[0], "adalah byte binary char index ke 0 dari string ", helo)
	fmt.Println(len(helo), "adalah panjang dari string ", helo)
	var num int32 = 10
	fmt.Println(num, "adalah integer")
	var num2 int64 = 1000000000000000000
	fmt.Println(num2, "adalah integer 64 bit")
	var betul bool = true
	fmt.Println(betul, "adalah boolean")
	var salah bool = false
	fmt.Println(salah, "adalah boolean")
}
