package main

import (
	"flag"
	"fmt"
)

func main() {
	//flag package is used to parse command line arguments
	//the first argument is the flag name, the second is the default value, and the third is the description
	playername := flag.String("player", "John Doe", "Player name")
	playerhp := flag.Int("hp", 100, "Player HP")
	playerlevel := flag.Int("level", 1, "Player level")

	flag.Parse()
	//if we run this program with the following command:
	//go run 04-flag.go -player "edward" -hp 120 -level 4
	//the output will be:
	//Player Name: edward
	//Player HP: 120
	//Player Level: 4
	//if we not specify the flag, it will use the default value
	fmt.Println("Player Name:", *playername)
	fmt.Println("Player HP:", *playerhp)
	fmt.Println("Player Level:", *playerlevel)
}
