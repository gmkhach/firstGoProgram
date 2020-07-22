package main

import "fmt"

var a = "This is a global variable. Use \"var\" for decleration of variables outside main."

func main() {
	fmt.Println("Hello everyone, this is the most awesome class in the entire world. We are having fun and learning the GO programming language.")

	b := "This vairable cannot be used outside of main"
	fmt.Println(b)

	// Delaring a variable without initializing it gives it a default value:
	// false - booleans
	// 0 - integers
	// 0.0 - floats
	// "" - strings
	// nil - pointers, functions, interfaces, slices, channels, maps
	var c bool
	var d int
	var e float32
	var f string
	var g interface{}
	fmt.Println("default boolean:", c)
	fmt.Println("default int:", d)
	fmt.Println("default float:", e)
	fmt.Println("default string:", f)
	fmt.Println("default interface:", g)

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
	// Declares a variable and assigns a value of a certain type
	x := 42
	fmt.Println("x is defined and given a value of", x)
	// Variable reassignment
	x = 99
	fmt.Println("x is reassigned to the value of", x)

	var y = "James Bond"
	fmt.Println("Bond,", y)
	
	// Statements 
	z := 100 + 24
	fmt.Println(z)
}

func foo() {
	fmt.Println("I'm in foo")
	fmt.Println(a)
}

func bar() {
	fmt.Println("and then we exited")
}
// Control flow:
// (1) sequence
// (2) loop; iterative
// (3) conditional
