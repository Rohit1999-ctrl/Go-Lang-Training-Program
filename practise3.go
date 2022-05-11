// Methods

// Create a method that sums two numbers
// Create a method that calls another method.
// Defer

// Predict what this code does:
// defer fmt.Println("Hello")
// defer fmt.Println("!")
// fmt.Println("World")
// Multiple return

// Change the return values from 2,4 to “hello”,”world”. Does it still work?
// Can a combination of strings and numbers be used?
// Variadic functions

// Create a variadic function that prints the names of students
// Recursion

// When is a function recursive?
// Can a recursive function call non-recursive functions?

package main

import (
	"fmt"
)

func sum(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}

func method1() {
	fmt.Println("first method called")
	method2()
}
func method2() {
	fmt.Println("second methods called")
}

func main() {

	sum(3, 5)
	method1()

}
