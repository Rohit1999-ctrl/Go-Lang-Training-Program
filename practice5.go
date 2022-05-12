//hamming distance

package main

import "fmt"

func Distance(a, b string) int {
	var count int
	fmt.Println("fsdf")
	for i := 0; i < len(a); i++ {
		if string([]rune(a)[i]) != string([]rune(b)[i]) {
			fmt.Println("", string([]rune(a)[i]), "=", string([]rune(a)[i]))
			count++
		}
	}
	return count

}

func main() {

	length := Distance("GAGCCTAACGGAT", "CARCGTAATGACG")
	fmt.Println("The length is: ", length)

}
