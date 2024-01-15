package main

import "fmt"

//by default, go lang variable is passed by value, not by reference

type Player struct {
	name   string
	hp, mp int
}

func main() {
	edward := Player{"Edward", 100, 20}
	edwardClone := edward //this will not pass the reference of edward

	edwardClone.name = "Edward Clone"
	edwardClone.hp -= 10
	edwardClone.mp -= 5

	//edward hp and mp is not changed
	//because edwardClone is a copy of edward
	fmt.Println(edward)
	fmt.Println(edwardClone)

	//to pass the reference of edward, we can use pointer
	//pointer is a variable that stores the address of another variable
	edwardPointer := &edward //same as var edwardPointer *Player = &edward
	edwardPointer.name = "Edward is changed"
	edwardPointer.hp += 10
	edwardPointer.mp += 5
	fmt.Println(edward) //edward hp and mp is changed

}
