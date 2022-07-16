package main

import "fmt"

func main() {
	forLoops()
	ifStatements()
	switchCaseStatement()
}

func forLoops() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
	
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}
}

func ifStatements() {
	var input int64
	fmt.Scanln("%d", &input)

	if input % 2 == 0 {
		fmt.Println("Number is even.")
	} else {
		fmt.Println("Number is odd.")
	}
}

func switchCaseStatement() {
	var i int64
	fmt.Println("Choose a number between 0 and 5: ")
	fmt.Scanln("%d", &i)
	switch i {
	case 0: fmt.Println("Zero")
	case 1: fmt.Println("One")
	case 2: fmt.Println("Two")
	case 3: fmt.Println("Three")
	case 4: fmt.Println("Four")
	case 5: fmt.Println("Five")
	default: fmt.Println("Unknown number")
	}
}