package main

import (
	"fmt"
	"os"
)

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
	fmt.Println(returnMultipleValues(xs))
	// Special functions accepting multiple parameters.
	fmt.Println(variadicFunction(int(xs[0]), int(xs[1]), int(xs[2])))
	xi := []int {2,3,5}
	// Similar to spread operator.
	fmt.Println(variadicFunction(xi...))
	closure()

	y := createEvenGenerator()
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(factorial(8))

	deferOrder()
	panicAndRecover()
}

// Average of an array of floating point numbers.
func average(xs []float64) float64 {
	total := 0.0

	for _, val := range xs {
		total += val
	}

	return total / float64(len(xs))
}

// A function can return multiple values.
func returnMultipleValues(xs []float64) (float64, float64) {
	x := make([] float64, 2)
	copy(x, xs)

	return x[0], x[1]
}

// Special function with multiple spreaded parameters.
func variadicFunction(args ...int) int {
	total := 0
	for _, val := range args {
		total += val
	}
	return total
}

// Special function that returns a function
// and remembers its most outer scope during next executions.
func closure() {
	x := 0

	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment) 		// Pointer
	fmt.Println(increment()) 	// Executed value
}

// Closure example.
func createEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// Recursive function example
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x - 1)
}

// Deferring specified function call to occur AFTER
// the second call.
func deferOrder() {
	x := func() {
		fmt.Println("X")
	}
	y := func() {
		fmt.Println("Y")
	}

	defer y()
	x()
}

// Defer can be used for resource freeups.
func fsWithCleanup() {
	f, _ := os.Open("index.html")
	defer f.Close()
}

// When paired with a defer function,
// panic recoveries become usable.
func panicAndRecover() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIK!")
}