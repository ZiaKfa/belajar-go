package main

import "fmt"

func main() {
	type hello string
	var world hello = "Hello World"
	fmt.Println(world)
}
