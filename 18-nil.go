package main

import "fmt"

//nil is a special value in Go that can be assigned to any type of variable
//nil is the default value for pointers, slices, maps, channels, and functions

func CreatePerson(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	john := CreatePerson("John Doe")
	if john == nil {
		fmt.Println("Name is nil")
	} else {
		fmt.Println(john)
	}
	empty := CreatePerson("")
	if empty == nil {
		fmt.Println("Name is nil")
	} else {
		fmt.Println(empty)
	}
}
