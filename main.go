package main

import (
	"fmt"
	"log"

	meow "github.com/abyanmajid/meow/core"
)

func main() {
	ageSchema := meow.Literal("Age", 123)
	result := ageSchema.Parse("Hello World!")
	if result.Error != nil {
		log.Fatalf(result.Error.Error())
	}

	fmt.Printf("Path: %v\nValue: %v\n", result.Path, result.Value)
}
