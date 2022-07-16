package main

import "fmt"

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // Still 5.

	// & symbol returns memory location of a variable.
	zeroP(&x)
	fmt.Println(x) // Changed.

	// 'new' is a builtin function allocates memory
	// for a variable.
	xPtr := new(int)
	zeroP(xPtr)
	fmt.Println(*xPtr)

}

func zero(x int) {
	x = 0
}

// * symbol dereferences the pointer variable,
// granting access to the value stored in memory.
func zeroP(xPtr *int) {
	*xPtr = 0
}