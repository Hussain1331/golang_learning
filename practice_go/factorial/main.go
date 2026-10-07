package main

import "fmt"

func main() {
	var num, factorial int
	factorial = 1
	fmt.Println("ENTER THE NUMBER TO FIND A FACTORIAL :-")
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		factorial = factorial * i
	}
	fmt.Println("THE FACTORIAL OF THE NUBER", num, "IS", factorial)
}
