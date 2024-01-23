package main

import "fmt"

//by default, go lang variable is passed by value, not by reference

type Player struct {
	name   string
	hp, mp int
}

func main() {
	edward := Player{"Edward", 100, 20}
	playerEditor := &edward
	playerEditor.name = "Edward is changed"

	//playerEditor = &Player{"Henry", 70, 30} this will not change edward, because the pointer will assigned to new Player struct
	*playerEditor = Player{"Henry", 70, 30}

	fmt.Println(edward)
	//fmt.Println(playerEditor)

	//using new keyword
	var cloud = new(Player) //this will create pointer to empty Player struct
	cloud.name = "Cloud"
	cloud.hp = 100
	cloud.mp = 50

	fmt.Println(cloud)

}
