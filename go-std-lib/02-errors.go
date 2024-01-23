package main

import "errors"

var (
	ValidationError = errors.New("Validation error")
	DatabaseError   = errors.New("Database error")
	NotFoundError   = errors.New("Not found error")
)

func StatusCode(id string) error {
	if id == "" {
		return ValidationError
	}

	if id == "404" {
		return NotFoundError
	}

	if id == "500" {
		return DatabaseError
	}

	return nil
}

func main() {
	err := StatusCode("8080")
	if err != nil {
		//old way
		// switch err {
		// case ValidationError:
		// 	println("Validation error")
		// case DatabaseError:
		// 	println("Database error")
		// case NotFoundError:
		// 	println("Not found error")
		// default:
		// 	println("Unknown error")
		// }
		if errors.Is(err, ValidationError) {
			println("Validation error")
		} else if errors.Is(err, DatabaseError) {
			println("Database error")
		} else if errors.Is(err, NotFoundError) {
			println("Not found error")
		} else {
			println("Unknown error")
		}
	} else {
		println("Running...")
	}
}
