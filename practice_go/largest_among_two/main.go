package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("ENTER THE FIRST NUMBER :-")
	fmt.Scanln(&num1)
	fmt.Println("ENTER NUMBER SECOND :-")
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Printf("NUMBER %d IS GREATER THEN NUMBER %d ", num1, num2)
	} else {
		fmt.Printf("NUMBER %d IS GREATER THEN NUMBER %d", num2, num1)

	}
}
