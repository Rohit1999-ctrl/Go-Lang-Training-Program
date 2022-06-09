//Given a number determine whether or not it is valid per the Luhn formula.
//The Luhn algorithm is a simple checksum formula used to validate a variety of identification numbers,
// such as credit card numbers and Canadian Social Insurance Numbers.
//The task is to check if a given string is valid.

package main

import (
	"fmt"
	"strconv"
)

func Valid(id string) bool {
	var sum int
	for i := 0; i < len(id); i++ {
		if i%2 == 0 {
			c := string([]rune(id)[i])
			v, err := strconv.Atoi(c)
			fmt.Println(v)
			if err == nil {
				l := v + v
				sum += l
			}
		} else {
			continue
		}
	}
	fmt.Println(sum)
	if sum%10 != 0 {
		return false
	} else {
		return true
	}
}

func main() {

	fmt.Printf("enter a number: ")
	var num string
	fmt.Scan(&num)
	l := Valid(num)
	fmt.Println(l)

}
