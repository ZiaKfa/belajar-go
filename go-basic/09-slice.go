package main

import "fmt"

/*slice is a piece of an array
just like array, slice is also a collection of data
but unlike array, slice is dynamic
slice is a reference type so slice is a pointer to an array
slice have 3 properties:
1. pointer to an array
2. length of slice (number of elements in slice, must be less than or equal to capacity)
3. capacity of slice (maximum length of slice)
*/

func main() {
	var fruit = [...]string{
		"apple", "orange", "mango", "papaya", "mangosteen", "lychee", "banana",
	}
	var slice1 = fruit[2:5]
	//slice1 = fruit[2:5] means slice1 is a slice of array fruit from index 2 to index 5
	var slice2 = fruit[2:]
	//slice2 = fruit[2:] means slice2 is a slice of array fruit from index 2 to the end of array
	var slice3 = fruit[:5]
	//slice3 = fruit[:5] means slice3 is a slice of array fruit from index 0 to index 5
	var slice4 = fruit[:]
	//slice4 = fruit[:] means slice4 is a slice of array fruit from index 0 to the end of array
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
	fmt.Println("slice3 = ", slice3)
	fmt.Println("slice4 = ", slice4)
	//length of slice
	fmt.Println("length of slice1 = ", len(slice1))
	fmt.Println("length of slice2 = ", len(slice2))
	fmt.Println("length of slice3 = ", len(slice3))
	fmt.Println("length of slice4 = ", len(slice4))
	//capacity of slice
	fmt.Println("capacity of slice1 = ", cap(slice1))
	fmt.Println("capacity of slice2 = ", cap(slice2))
	fmt.Println("capacity of slice3 = ", cap(slice3))
	fmt.Println("capacity of slice4 = ", cap(slice4))
	//slice is a reference type, so if we change the value of slice, the value of array will also change
	slice1[0] = "grape"
	slice1[2] = "watermelon"
	slice2[1] = "melon"
	slice3[1] = "pineapple"
	slice4[5] = "durian"
	slice4[6] = "jackfruit"
	fmt.Println("changed array = ", fruit)

	//append slice
	//append slice is a function to add element to slice, if the length of slice is greater than
	//the capacity of slice, a new array will be created and the value of slice will be copied to the new array
	var slice5 = fruit[2:5]
	var slice6 = append(slice5, "cucumber", "tomato", "grape", "strawberry")
	fmt.Println("slice5 = ", slice5)
	fmt.Println("appended slice = ", slice6)
	slice6[0] = "cherry"
	slice6[1] = "lemon"
	slice6[2] = "lime"
	slice6[3] = "date"
	fmt.Println("changed appended slice = ", slice6) //slice6 changed but slice5 and fruit didn't change
	fmt.Println("array = ", fruit)

	//make slice
	//make slice is a function to create a slice
	//make slice have 3 parameters:
	//1. type of slice
	//2. length of slice
	//3. capacity of slice
	var slice7 = make([]string, 3, 5) //this will create a slice of string with length 3 and capacity 5, and also a new hidden array with length 5
	slice7[0] = "apple"
	slice7[1] = "banana"
	slice7[2] = "grape"
	fmt.Println("slice7 = ", slice7)
	fmt.Println("length of slice7 = ", len(slice7))
	fmt.Println("capacity of slice7 = ", cap(slice7))

	//copy slice
	//copy slice is a function to copy a slice to another slice
	//copy slice have 2 parameters:
	//1. destination slice
	//2. source slice
	var slice8 = make([]string, len(slice7), cap(slice7))
	copy(slice8, slice7)
	fmt.Println("slice8 = ", slice8)
	fmt.Println("length of slice8 = ", len(slice8))
	fmt.Println("capacity of slice8 = ", cap(slice8))

}
