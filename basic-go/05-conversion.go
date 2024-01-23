package main

import "fmt"

func main() {
	//number conversion
	var num32 int32 = 10000000
	var num64 int64 = int64(num32)
	//will cause int overflow cause the value is exceeding the limit for int16
	var num16 int16 = int16(num32)
	fmt.Println(num64)
	fmt.Println(num32)
	fmt.Println(num16, "this is an overflow, the value will reset to the minimum value of int16")

	//string conversion
	var hello string = "Hello World"
	var helloByte byte = hello[0]
	var helloString string = string(helloByte)
	fmt.Println(helloByte, "is a binary byte of ", helloString)
}
