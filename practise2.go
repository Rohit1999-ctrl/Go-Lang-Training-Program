// Struct

// Create a struct house with variables noRooms, price and city
// How does a struct differ from a class?
// Maps

// What is a map?
// Is a map ordered?
// What can you use a map for?
// Random numbers

// Make a program that rolls a dice (1 to 6)
// Can you generate negative numbers?
// Pointers

// Where are variables stored in the computer?
// What is a pointer?
// How can you declare a pointer?
// Slices

// Take the string ‘hello world’ and slice it in two.
// Can you take a slice of a slice?

package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type house struct {
	noRooms string
	price   int
	city    string
}

func main() {

	MyHouse := house{
		"yes", 1200, "new york",
	}
	fmt.Println(MyHouse.price)
	fmt.Println(MyHouse.noRooms)
	fmt.Println(MyHouse.city)

	fmt.Println(rand.Intn(6))

	name := "hello world"

	names1 := strings.Split(name, " ")
	fmt.Println(names1)
	fmt.Println(names1[0])

}
