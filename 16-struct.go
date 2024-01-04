package main

import "fmt"

type Person struct {
	Name, Address string
	Age           int
}

func main() {
	newPerson := Person{
		Name:    "John Doe",
		Address: "Indonesia",
		Age:     20,
	}
	fmt.Println(newPerson)

	somePerson := Person{"John Cena", "Indonesia", 21}
	fmt.Println(somePerson)

	var otherPerson Person
	otherPerson.Name = "Jane Doe"
	otherPerson.Address = "Indonesia"
	otherPerson.Age = 19
	fmt.Println(otherPerson.Name)
	fmt.Println(otherPerson.Address)
	fmt.Println(otherPerson.Age)

}
