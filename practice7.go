//chessboard
package main

import "fmt"

func Square(number int) int {
	return number * number
}

func Total() int {
	sum := 0
	for i := 0; i < 65; i++ {
		sum = sum + (i * i)
	}
	return sum
}

func main() {
	var n int
	fmt.Println("enter the block no: ")
	fmt.Scan(&n)
	s := Square(n)
	fmt.Println("The number of gains in the block will be: ", s)
	p := Total()
	fmt.Println("the total number of grains in the whole chessboard is: ", p)
}
