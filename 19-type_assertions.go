package main

import "fmt"

func Omg() interface{} {
	return true
}

func main() {
	omaga := Omg() // omaga is interface{} not string
	// fmt.Println(omaga)
	// omagaString := omaga.(string)
	// fmt.Println(omagaString)
	// omagaInt := omaga.(int)
	// fmt.Println(omagaInt) // panic: interface conversion: interface {} is string, not int because omaga is string
	switch value := omaga.(type) {
	case string:
		fmt.Println(value, "is string")
	case int:
		fmt.Println(value, "is int")
	case bool:
		fmt.Println(value, "is bool")
	default:
		fmt.Println(value, "is unknown")
	}
}
