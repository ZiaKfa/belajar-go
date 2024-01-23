package main

import (
	"fmt"
	"strings"
)

func main() {
	//strings.Contains() returns true if the string contains the substring
	fmt.Println(strings.Contains("Hello World", "Hello"))
	fmt.Println(strings.Contains("Hello World", "hello"))

	//strings.Count() returns the number of non-overlapping instances of substring
	fmt.Println(strings.Count("Hello World", "l"))

	//strings.HasPrefix() returns true if the string starts with the prefix
	fmt.Println(strings.HasPrefix("Hello World", "Hello"))
	fmt.Println(strings.HasPrefix("Hello World", "World"))

	//strings.HasSuffix() returns true if the string ends with the suffix
	fmt.Println(strings.HasSuffix("Hello World", "Hello"))
	fmt.Println(strings.HasSuffix("Hello World", "World"))

	//strings.Index() returns the index of the first instance of substring
	fmt.Println(strings.Index("Hello World", "World"))

	//strings.Join() returns a string by joining all the elements of a slice with the separator
	fmt.Println(strings.Join([]string{"Hello", "World"}, " "))

	//strings.Replace() returns a string by replacing the old string with the new string n times
	fmt.Println(strings.Replace("Hello World", "l", "x", 1))
	fmt.Println(strings.Replace("Hello World", "l", "x", 2))
	fmt.Println(strings.Replace("Hello World", "l", "x", -1))

	//strings.Split() returns a slice of strings by splitting the string with the separator
	fmt.Println(strings.Split("Hello World", " "))

	//strings.ToLower() returns a string by converting all characters to lower case
	fmt.Println(strings.ToLower("Hello World"))

	//strings.ToUpper() returns a string by converting all characters to upper case
	fmt.Println(strings.ToUpper("Hello World"))

	//strings.TrimSpace() returns a string by removing all leading and trailing white space
	fmt.Println(strings.TrimSpace(" Hello World "))

	//strings.Trim() returns a string by removing all leading and trailing characters
	fmt.Println(strings.Trim("Hello World", "Hd"))

	//strings.TrimLeft() returns a string by removing all leading characters
	fmt.Println(strings.TrimLeft("Hello World", "Hd"))

	//strings.TrimRight() returns a string by removing all trailing characters
	fmt.Println(strings.TrimRight("Hello World", "Hd"))
}
