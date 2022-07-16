package main

import (
	"fmt"

	math "eray_math"
)

// import m "eray_math/math.go"
// Won't work in module mode.
func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
