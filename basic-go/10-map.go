package main

import "fmt"

func main() {
	/*Map is a collection of key-value pair, like dictionary in Python or associative array in PHP
	The difference is that the key and value in Go can be any data type, but must be the same for all elements
	Map is reference type, so it must be initialized using make() function
	The key must be unique, and the value can be duplicated
	the
	*/
	var bookPrice = map[string]int{
		"Math":    100,
		"English": 90,
		"Science": 95,
	}
	fmt.Println(bookPrice)
	fmt.Println(bookPrice["Math"])

	//change the value
	bookPrice["Math"] = 120
	fmt.Println("changed math book price to ", bookPrice["Math"])

	//add new element
	bookPrice["History"] = 80
	fmt.Println("added history book to the map : ", bookPrice)

	//delete element
	delete(bookPrice, "History")
	fmt.Println("deleted history book to the map : ", bookPrice)

	//map length
	fmt.Println("map length : ", len(bookPrice))

	//use this if you want to make empty map and add the element later
	var person map[string]string = make(map[string]string)
	person["name"] = "John Doe"
	person["age"] = "20"
	person["address"] = "Indonesia"
	person["hobby"] = "Coding"
	person["job"] = "Programmer"
	fmt.Println(person)

}
