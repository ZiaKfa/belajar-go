package main

import "fmt"

func main() {
	//manual array declaration
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4

	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
	fmt.Println(arr[3])
	//length of array
	var length = len(arr)
	fmt.Println("the length of the array is ", length)

	//direct array declaration
	var names = [4]string{
		"John", "Doe", "Will", "Smith",
	}
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])
	//length of array
	var length2 = len(names)
	fmt.Println("the length of the array is ", length2)
}
