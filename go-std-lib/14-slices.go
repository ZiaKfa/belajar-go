package main

import (
	"fmt"
	"slices"
)

func main() {
	pokemons := []string{"Pikachu", "Bulbasaur", "Charmander"}
	number := []int{1, 2, 3}

	fmt.Println(slices.Max(number))
	fmt.Println(slices.Min(number))
	fmt.Println(slices.Max(pokemons))
	fmt.Println(slices.Min(pokemons))
	fmt.Println(slices.Contains(pokemons, "Pikachu"))
	fmt.Println(slices.Contains(pokemons, "Mewtwo"))
	fmt.Println(slices.Index(pokemons, "Pikachu"))
}
