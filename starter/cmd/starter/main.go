package main

import (
	"fmt"
	"os"
)

func main() {
	name := "World"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Printf("Hello, %s! 🚀\n", name)
	fmt.Println("Bazel + Go is working!")
}

