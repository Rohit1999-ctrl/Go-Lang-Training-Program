package main

import "fmt"

func IsIsogram(word string) bool {
	check := 0
	for i := 0; i < len(word); i++ {
		count := 0
		for j := 0; j < len(word); j++ {
			fmt.Println(" ", string([]rune(word)[i]), " = ", string([]rune(word)[j]))
			if string([]rune(word)[i]) == string([]rune(word)[j]) {
				count++
			}
		}
		fmt.Println("count = ", count)
		if count < 1 {
			check = 1
		} else if check <= 2 {
			check = 1
		}
		fmt.Println("check=", check)

	}
	if check == 1 {
		return false
	} else {
		return true
	}

}

func main() {
	var word string
	fmt.Scan(&word)
	v := IsIsogram(word)
	fmt.Println(v)
}
