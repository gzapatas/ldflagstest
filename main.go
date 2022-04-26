package main

import (
	"fmt"

	"github.com/gzapatas/ldflagstest/build"
)

var Version string = "deployment"

func Suma(a, b int) int {
	return a + b
}

func main() {
	fmt.Printf("Version = %v\n", Version)
	fmt.Printf("Time = %v\n", build.Time)
	fmt.Printf("User = %v\n", build.User)

	fmt.Printf("Suma = %v\n", Suma(2, 5))
}
