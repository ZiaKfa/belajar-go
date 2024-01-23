package main

import (
	"fmt"
	"math"
)

func main() {
	value := 3.14
	fmt.Println(math.Floor(value))
	fmt.Println(math.Ceil(value))
	fmt.Println(math.Round(value))
	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Min(1, 2))
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Sqrt(16))
	fmt.Println(math.Abs(-4))
	fmt.Println(math.Sin(0))
	fmt.Println(math.Cos(0))
	fmt.Println(math.Tan(0))
}
