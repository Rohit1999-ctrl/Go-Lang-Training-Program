//scramble

package main

import "fmt"

func Score(word string) int {
	points := 0
	for i := 0; i < len(word); i++ {
		x := string([]rune(word)[i])
		if x == "A" || x == "E" || x == "I" || x == "O" || x == "U" || x == "L" || x == "N" || x == "R" || x == "S" || x == "T" {
			points = points + 1
		} else if x == "D" || x == "G" {
			points = points + 2
		} else if x == "B" || x == "C" || x == "M" || x == "P" {
			points = points + 3
		} else if x == "F" || x == "H" || x == "V" || x == "W" || x == "Y" {
			points = points + 4
		} else if x == "K" {
			points = points + 5
		} else if x == "J" || x == "X" {
			points = points + 8
		} else if x == "Q" || x == "Z" {
			points = points + 10
		}

	}
	return points
}

func main() {
	var word string
	fmt.Scan(&word)
	total := Score(word)
	fmt.Println(total)
}
