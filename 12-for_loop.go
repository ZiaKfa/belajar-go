package main

import "fmt"

/*unlike other programming language, Go only has for loop,
  there is no while loop or do-while loop in Go, but you can use for loop to do that
*/

func main() {
	//for loop without condition
	max := 10
	count := 1
	for count <= max {
		fmt.Println("loop", count)
		count++
	}
	//in other programming language, for loop must have 3 statements : initialization, condition, and increment/decrement
	//but in Go, you can use for loop like while loop like the code above

	//for loop with condition
	for i := 10; i >= 1; i-- {
		fmt.Println("loop", i)
	}

	//looping through array, slice, or map
	fruits := []string{"apple", "banana", "orange", "watermelon", "melon"}
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	//looping through array, slice, or map with range
	for index, fruit := range fruits {
		fmt.Println(index, fruit)
	}

	//if the index is not used, you can replace it with underscore (_)
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	//looping through map
	person := map[string]string{
		"name":    "John Doe",
		"address": "Indonesia",
		"age":     "20",
		"hobby":   "Coding",
	}
	//in map, the range can only return key and value, unlike array or slice that can return index and value
	for key, value := range person {
		fmt.Println(key, value)
	}

	//break and continue
	for i := 1; i <= 10; i++ {
		if i == 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}
	//this will stop the loop when i == 5
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	//this will skip the even number, and continue the loop on odd number
}
