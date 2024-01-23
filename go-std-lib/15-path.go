package main

import (
	"fmt"
	"path"          //used for path manipulation
	"path/filepath" //used for os independent path manipulation
)

func main() {
	fmt.Println(path.Dir("/home/username/14148.txt"))
	fmt.Println(path.Base("/home/username/14148.txt"))
	fmt.Println(path.Ext("/home/username/14148.txt"))
	fmt.Println(path.Join("home", "username", "14148.txt"))

	fmt.Println(filepath.Dir("/home/username/14148.txt"))
	fmt.Println(filepath.Base("/home/username/14148.txt"))
	fmt.Println(filepath.Ext("/home/username/14148.txt"))
	fmt.Println(filepath.IsAbs("/home/username/14148.txt"))
	fmt.Println(filepath.IsLocal("/home/username/14148.txt"))
	fmt.Println(filepath.Join("home", "username", "14148.txt"))

}
