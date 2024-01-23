package main

import (
	"LearnGO/trypackage"
	"fmt"
)

func main() {
	fmt.Println(trypackage.SayHello("John Doe"))
	fmt.Println(trypackage.PublicVariable)
	//fmt.Println(trypackage.privateVariable) //this will error because privateVariable is private
	//fmt.Println(trypackage.privateFunction()) //this will error because privateFunction is private
}
