package main

import "fmt"

func main() {
	var num, first_digit int
	fmt.Println("ENTER THE NUMBER TO FIND FIRST DIGIT :-")
	fmt.Scanln(&num)
	for num > 10 {
		num = num / 10
		first_digit = num
	}
	fmt.Println(first_digit)
}
