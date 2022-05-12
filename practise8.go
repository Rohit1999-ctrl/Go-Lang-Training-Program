//raindrop

package main

import "fmt"

func Convert(number int) string {
	word := ""
	if number%3 == 0 && number%5 == 0 && number%7 == 0 {
		word += "PLingPlangPlong"
	} else if number%3 == 0 && number%5 == 0 {
		word += "PlingPlang"
	} else if number%3 == 0 && number%7 == 0 {
		word += "PlingPlong"
	} else if number%5 == 0 && number%7 == 0 {
		word += "PlangPlong"
	} else if number%3 == 0 {
		word += "Pling"
	} else if number%5 == 0 {
		word += "Plang"
	} else if number%7 == 0 {
		word += "Plong"
	}
	return word
}

func main() {
	var num int
	fmt.Println("enter a number: ")
	fmt.Scan(&num)
	w := Convert(num)
	fmt.Println(w)
}
