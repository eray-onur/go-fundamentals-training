package main

import "fmt"

func main() {
	// arrayBasics()
	// arrayOperations()
	//mapBasics()
	nestedMaps()
}

func nestedMaps() {
	elements := map[string] map[string]string {
		"H": {
			"name": "Hydrogen",
			"state": "gas",
		},
	}

	if el, ok := elements["H"]; ok {
		fmt.Println("Element by the name %s is found.", el["name"])
	}
}

func mapBasics() {
	x := make(map[string] int)
	x["key"] = 10
	fmt.Println(x)
	// deleting a key value from map.
	delete(x, "key")
	fmt.Println(x)
	x["key"] = 20

	// 1st: value of given key on map.
	// 2nd: whether if lookup was successful or not.
	name, ok := x["key"]

	// IF lookup is successful, continue to inner scope.
	if name, ok := x["key"]; ok {
		fmt.Println(name, ok)
	}

	fmt.Println(name, ok)
}

func arrayOperations() {
	arr := [5]float64{1, 2, 3, 4, 5}
	// x : y => x: starting index, y: item count
	x := arr[0:3] 

	fmt.Println(x)
	
	// Appending an item at the end of array.

	appendedArray := append(x, 4, 5, 6)
	fmt.Println("Newly appended array:")
	fmt.Println(appendedArray)

	// Copying items into the destination array.
	srcArray := []int{1, 2, 3}
	destArray := make([] int, 2)
	copy(destArray, srcArray)
	fmt.Println("Copied array:")
	fmt.Println(destArray)
}

func arrayBasics() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	var total float64 = 0
	var float_array [5]float64

	float_array[0] = 55
	float_array[3] = 20

	// Regular for loop.
	for i:=0; i < 5; i++ {
		total += float_array[i]
	}
	fmt.Println(total)
	fmt.Println("Average total:")
	fmt.Println(total / float64(len(x)))

	total = 0

	// Go compiler never allows unused variables to
	// persist, which is why a dash ('_') is necessary.
	// Equivalent of 'for each' loops in many languages.
	for _ , value := range float_array {
		total += value
	}
	fmt.Println(total / float64(len(float_array)))

}