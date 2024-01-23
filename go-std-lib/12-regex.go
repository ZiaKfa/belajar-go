package main

import (
	"fmt"
	"regexp"
)

//use RE2 syntax

func main() {
	regex := regexp.MustCompile(`[a-z]+`)
	fmt.Println(regex.MatchString("halo"))
	fmt.Println(regex.MatchString("123"))

	fmt.Println(regex.FindAllString("halo halo bandung 123", 10))
}
