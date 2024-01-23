package main

import (
	"container/ring"
	"fmt"
)

func main() {
	//container/ring is a circular list implementation in Go
	dataRing := ring.New(5)
	for i := 0; i < dataRing.Len(); i++ {
		dataRing.Value = i
		if i == 0 {
			dataRing.Value = "Pikachu"
		} else if i == 1 {
			dataRing.Value = "Charmander"
		} else if i == 2 {
			dataRing.Value = "Squirtle"
		} else if i == 3 {
			dataRing.Value = "Bulbasaur"
		} else if i == 4 {
			dataRing.Value = "Mewtwo"
		}
		dataRing = dataRing.Next()
	}

	dataRing.Do(func(data interface{}) {
		fmt.Println(data)
	})
}
