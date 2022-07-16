package main

import (
	"flag"
	"fmt"
	"math/rand"
)

// run the program with -v for flag to work
// ex. (go run -v 12_command_line_args.go --max 100)
func main() {
	// Define flags
	maxp := flag.Int("max", 6, "the max value")
	// Parse
	flag.Parse()
	// Generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
}