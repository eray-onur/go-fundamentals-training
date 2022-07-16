package main

import (
	"errors"
	"fmt"
)

func main() {
	// Creating custom error.
	err := errors.New("error message")
	fmt.Println(err)
}