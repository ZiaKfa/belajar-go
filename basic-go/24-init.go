package main

import (
	"LearnGO/database"
	_ "LearnGO/initial" // _ is used to ignore the package, but still execute the init function
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
