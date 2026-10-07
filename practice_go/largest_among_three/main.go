package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Println("ENTER THE FIRST NUMBER :-")
	fmt.Scanln(&num1)
	fmt.Println("ENTER NUMBER SECOND :-")
	fmt.Scan(&num2)
	fmt.Println("ENTER NUMBER THIRD:-")
	fmt.Scan(&num3)

	if num1 > num2 && num1 > num3 {
		fmt.Printf("%d is largest among 3", num1)

	} else if num2 > num1 && num2 > num3 {
		fmt.Printf("%d is the largest number among 3", num2)
	} else {
		fmt.Printf("%d is the largest number among 3", num3)
	}
}
