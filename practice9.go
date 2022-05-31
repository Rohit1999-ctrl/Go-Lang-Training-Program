//error detection
//to detect if the first number is gretaer than second number, and if it is give a user defined error
package main

import (
	"errors"
	"fmt"
)

func main() {
	var f int
	var l int
	fmt.Println("enter the first number: ")
	fmt.Scan(&f)
	fmt.Println("enter the last number: ")
	fmt.Scan(&l)
	//fmt.Println(addition(f, l))
	check, err := addition(f, l)
	if err != nil {
		fmt.Println("an error was found", err)
	} else {
		fmt.Println(check)
	}
}

func addition(first, last int) (int, error) {
	if first > last {
		return 0, errors.New("first number is greater than last")
	} else {
		total := 0
		for i := first; i <= last; i++ {
			total += i
		}
		return total, nil
	}

}
