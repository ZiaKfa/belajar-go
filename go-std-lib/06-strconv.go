package main

import (
	"fmt"
	"strconv"
)

func main() {
	//strconv package is used to convert string to other types
	rightThing, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(rightThing)
	}
	resultInt, err := strconv.Atoi("100") //Atoi is short for "ASCII to integer"
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resultInt)
	}

	//other than ParseBool and Atoi, strconv also has ParseInt, ParseUint, ParseFloat, FormatBool, FormatInt, FormatUint, FormatFloat, and many more
	//it can also convert numeric types, for example:
	binary := strconv.FormatInt(100, 2) //will return "1100100"
	fmt.Println(binary)

	octal := strconv.FormatInt(100, 8) //will return "144"
	fmt.Println(octal)

	hex := strconv.FormatInt(31, 16) //will return "1f"
	fmt.Println(hex)

}
