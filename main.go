package main

import (
	"fmt"
	"os"

	"github.com/squeakycheese75/panicgen/generator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: panicgen <go-source-dir>")
		os.Exit(1)
	}

	dir := os.Args[1]
	err := generator.GenerateTests(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "panicgen failed: %v\n", err)
		os.Exit(1)
	}
}
