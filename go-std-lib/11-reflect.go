package main

import (
	"fmt"
	"reflect"
)

type Pokemon struct {
	Name string `type:"electric" weak:"ground"`
	HP   int    `max:"100"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	var valueValue reflect.Value = reflect.ValueOf(value)

	fmt.Println("Type:", valueType)
	fmt.Println("Value:", valueValue)
	fmt.Println("TypeName:", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		fmt.Println("Field", i, ":", valueType.Field(i).Name+":", valueValue.Field(i).Type())
		value := valueType.Field(i)
		fmt.Println("Tag Type:", value.Tag.Get("type"))
		fmt.Println("Tag Weak:", value.Tag.Get("weak"))
		fmt.Println("Tag Max:", value.Tag.Get("max"))
	}
}

func isElectric(value any) bool {
	var valueType reflect.Type = reflect.TypeOf(value)
	for i := 0; i < valueType.NumField(); i++ {
		if valueType.Field(i).Tag.Get("type") == "electric" {
			return true
		}
	}
	return false
}

func main() {
	readField(Pokemon{"Pikachu", 100})
}
