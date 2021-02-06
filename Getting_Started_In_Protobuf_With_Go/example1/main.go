package main

import (
	"fmt"
)

func main() {
	var person Person

	person.Id = 1001
	person.Email = "abc@xyz.in"
	person.IsActive = true

	fmt.Println(&person)
}
