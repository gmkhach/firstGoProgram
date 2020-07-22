package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, this is the most awesome class in the entire world. We are having fun and learning the GO programming language.")
	
	foo()
	
	fmt.Println("starting a loop")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// Use the "_" character to omitt any returned values you do not intend to use
	// In GO you have to use all of the declared variables. "_" allows to ignore values you do not intend to use.
	n, _ := fmt.Println("Hello, playground", 42, true)
	fmt.Println(n)

	bar()
	
	// Short declaration operator
	x := 42
	fmt.Println("x is defined and given a value of", x)
	// Variable reassignment
	x = 99
	fmt.Println("x is reassigned to the value of", x)

	y := "James Bond"
	fmt.Println("Bond,", y)
	
	// Statements 
	z := 100 + 24
	fmt.Println(z)
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("and then we exited")
}
// Control flow:
// (1) sequence
// (2) loop; iterative
// (3) conditional
