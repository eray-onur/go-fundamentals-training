package main

import "fmt"

func main() {
	var x string = "Hello, world!"
	fmt.Println(x)
	var y string
	y = "Lorem"
	fmt.Println(y)
	z := "bb"
	fmt.Println(z)
	print("cc")

	const t string = "Lorem ipsum"
	print(t)
}

func print(x string) {
	fmt.Println(x)
}
