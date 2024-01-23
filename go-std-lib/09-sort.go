package main

import (
	"fmt"
	"sort"
)

func main() {
	//sort is a package that implements sorting for builtins and user-defined types
	data := []string{"Pikachu", "Charmander", "Squirtle", "Bulbasaur", "Mewtwo"}
	sort.Strings(data)
	fmt.Println(data)

	dataInt := []int{4, 3, 1, 5, 2}
	sort.Ints(dataInt)
	fmt.Println(dataInt)

	//sort.Sort() can be used to sort user-defined types
	type Pokemon struct {
		Name string
		HP   int
	}

	pokemons := []Pokemon{
		{"Pikachu", 100},
		{"Charmander", 120},
		{"Squirtle", 90},
		{"Bulbasaur", 200},
		{"Mewtwo", 170},
	}

	sort.Slice(pokemons, func(i, j int) bool {
		return pokemons[i].HP < pokemons[j].HP
	})

	fmt.Println(pokemons)
}
