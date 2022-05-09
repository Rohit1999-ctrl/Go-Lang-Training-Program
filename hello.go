package main

import (
	"fmt"
)

//link of the exersice: https://golangr.com/exercises/
// Hello world

// Create a program that shows your name
// Create a program that shows your address
// Comments

// Create a program and add a comment with your name
// Strings

// Create a program with multiple string variables
// Create a program that holds your name in a string.
// Keyboard input

// Make a program that lets the user input a name
// Get a number from the console and check if it’s between 1 and 10.
// Variables

// Calculate the year given the date of birth and age
// Create a program that calculates the average weight of 5 people.
// Scope

// What’s the difference between a local and global variable?
// How can you make a global variable?
// Arrays

// Create an array with the number 0 to 10
// Create an array of strings with names
// For loops

// Can for loops exist inside for loops?
// Make a program that counts from 1 to 10.
// Range

// What is the purpose of range ?
// What the difference between the line for index, element := range a and the line for _, element := range a ?
// If statements

// Make a program that divides x by 2 if it’s greater than 0
// Find out if if-statements can be used inside if-statements.
// While loops

// How does a while loop differ from a for loop?

func main() {

	// fmt.Println("rohit")
	// fmt.Print("biju nowas, bepquegal goa")
	// name := "rohit"
	// address := "biju niwas, bepquegal goa"
	// fmt.Println(name, address)

	// fmt.Print("enter your name: ")
	// var first string
	// fmt.Scan(&first)
	// fmt.Println(first)
	// fmt.Println("enter a number: ")
	// var number int
	// fmt.Scan(&number)
	// if number > 1 && number < 10 {
	// 	fmt.Println("yes, finally")
	// } else {
	// 	fmt.Println("no man")
	// }
	// var year int
	// fmt.Println("enter your age: ")
	// fmt.Scan(&year)
	// var check int = 2022 - year
	// fmt.Println("My year is: ", check)

	// var sum int

	// for i := 0; i < 5; i++ {
	// 	var w int
	// 	fmt.Println("enter the weight: ")
	// 	fmt.Scan(&w)
	// 	sum += w
	// }
	// fmt.Print("the average weight is: ", sum/5)

	// names := [...]string{"rohit", "biju", "boy"}
	// fmt.Print(names)
	// for i := 1; i < 11; i++ {
	// 	fmt.Println(i)
	// }

	fmt.Print("enter a number: ")
	var num int
	fmt.Scan(&num)
	if num > 0 {
		fmt.Print("the sol is: ", num/2)
	}
}
