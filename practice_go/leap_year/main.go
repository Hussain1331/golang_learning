package main

import "fmt"

func main() {
	var year int
	fmt.Println("ENTER THE YEAR U WANNA CHECK78 :-")
	fmt.Scanln(&year)

	if year%4 == 0 {
		fmt.Println("the year is leap year")
	} else {
		fmt.Println("the year is not leap year")

	}
}
