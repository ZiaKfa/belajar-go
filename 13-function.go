package main

import "fmt"

//function without parameter and return value
func sayHello() {
	fmt.Println("Hello World")
}

//function with parameter and without return value
func printName(name string) {
	fmt.Println("Hello", name)
}

//function with parameter and return value
func getName(name string) string /*<== this is the return type*/ {
	return "Hello " + name
}

//function with multiple return value
//unlike other programming language, Go function can return multiple value
func getFullName(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

func getFullNameArray(firstName string, lastName string) [2]string {
	return [2]string{firstName, lastName}
}

//variadic function
//variadic function is a function that can accept multiple parameter
//the parameter used in variadic function is the last parameter

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//function as parameter
//in Go, function can be used as parameter
func holiday(day string, check func(string) bool) {
	if check(day) {
		fmt.Println("Today is Holiday")
	} else {
		fmt.Println("Today is", day)
	}
}

func isHoliday(day string) bool {
	if day == "Sunday" || day == "Saturday" {
		return true
	} else {
		return false
	}
}

//recursive function
func factorial(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * factorial(number-1)
	}
}

//we can use type declaration to make the code more readable
// type CheckDay func(string) bool

func main() {
	sayHello()
	printName("John Doe")
	fmt.Println(getName("John Doe"))
	/* the return variable here can be assigned to ( _ ) if the value is not used */
	firstName, lastName := getFullName("John", "Doe")
	fmt.Println(firstName, lastName)
	fullName := getFullNameArray("John", "Doe")
	fmt.Println(fullName[0])
	fmt.Println(fullName[1])

	//no matter how many parameter you pass to the function, it will be working just fine
	summary := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(summary)

	someNumbers := []int{1, 2, 3, 4, 5}
	total := sumAll(someNumbers...) //this is how you pass slice to variadic function
	fmt.Println(total)

	//function as parameter
	holiday("Sunday", isHoliday)
	holiday("Monday", isHoliday)

	//anonymous function
	//anonymous function is a function without name
	//anonymous function can be assigned to a variable
	holidayFunc := func(day string) bool {
		if day == "Sunday" || day == "Saturday" {
			return true
		} else {
			return false
		}
	}
	holiday("Sunday", holidayFunc)
	//is the same as
	holiday("Sunday", func(day string) bool {
		return day == "Sunday" || day == "Saturday"
	})

	//recursive function
	fmt.Println(factorial(5))
}
