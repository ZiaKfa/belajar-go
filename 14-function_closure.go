package main

import "fmt"

func main() {
	count := 0

	increment := func() {
		fmt.Println("add 1 to count")
		count++ //this count is from the main function
	}

	increment()
	increment()

	fmt.Println(count)
}
