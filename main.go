package main

import (
	"fmt"

	"github.com/gzapatas/ldflagstest/build"
)

var Version string = "deployment"

func main() {
	fmt.Printf("Version = %v\n", Version)
	fmt.Printf("Time = %v\n", build.Time)
	fmt.Printf("User = %v\n", build.User)
}
