package main

import "fmt"

func main() {
	var num int
	fmt.Println("ENTER THE NUMBER U WANT TABLE OF :-")
	fmt.Scanln(&num)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, (num * i))
	}
}
