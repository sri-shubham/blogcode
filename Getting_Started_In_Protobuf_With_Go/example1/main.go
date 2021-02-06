package main

import (
	"example1/pb"
	"fmt"
)

func main() {
	var person pb.Person

	person.Id = 1001
	person.Email = "abc@xyz.in"
	person.IsActive = true

	fmt.Println(&person)
}
