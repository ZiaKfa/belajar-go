package main

import (
	"container/list"
	"fmt"
)

func main() {
	//container/list is a doubly linked list implementation in Go
	dataList := list.New()
	dataList.PushBack("Pikachu")
	dataList.PushBack("Charmander")
	dataList.PushBack("Squirtle")
	dataList.PushBack("Bulbasaur")
	dataList.PushFront("Mewtwo")

	head := dataList.Front()
	fmt.Println(head.Value)

	for e := dataList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	for e := dataList.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
