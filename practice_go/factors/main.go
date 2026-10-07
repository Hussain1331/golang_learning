package main

import "fmt"

func main() {
	var num int
	fmt.Println("ENTER THE NUMBER TO FIND A FACTOR :-")
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println()
}
