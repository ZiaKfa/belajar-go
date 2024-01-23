package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return a / b, nil
}
func main() {
	result, err := divide(100, 0)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Result:", result)
	}
}
