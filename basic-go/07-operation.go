package main

import "fmt"

func main() {
	//math operation
	var a int = 10
	var b int = 5
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	//increment and decrement
	var i int = 0
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)

	//comparison operation
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a < b)
	fmt.Println(a <= b)

	//logical operation
	fmt.Println(a == 10 && b == 5)
	fmt.Println(a == 10 || b == 5)
	fmt.Println(!(a == 10 && b == 5))
}
