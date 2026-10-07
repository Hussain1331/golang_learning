// PRACTICE SET 1
package main

import (
	"fmt"
)

// func main() {
// 	fmt.Println("HELLO WORLD")
// }

// PROGRAM TO FIND THE SUM OF TWO NUMBERS _______________
// func main() {
// 	var num1 int = 2
// 	var num2 int = 3
// 	var marks float32 = 45.5
// 	fmt.Println("marks obtained ", marks)
// 	var sum int = num1 + num2
// 	fmt.Println("Sum of", num1, "and", num2, "is", sum)
// }

//_________________________________________________________
//PROGRAM TO FIND THE CIRCUMFERENCE OF CIRCLE
// //func main() {
// 	const pi float32 = 3.1415
// 	var radius float32 = 12.0
// 	arr2 := [5]int{4, 5, 6, 7, 8}
// 	fmt.Println(arr2)
// 	fmt.Println("CIRCUMFERENCE OF CIRCLE :", 2*pi*radius)
// }

//Go Program to Find the Cube of a Number
// func main() {
// 	var number int
// 	fmt.Println("ENTER NUMBER")
// 	fmt.Scanln(&number)
// 	fmt.Println("THE CUBE OF", number, "IS:", number*number*number)
// }

//_______________________________________
//program to find the number of digit

// func main() {
// 	var num, count int
// 	count = 0
// 	fmt.Println("ENTER ANY NUMBER YOU WANT TO COUNT")
// 	fmt.Scanln(&num)
// 	for num > 0 {
// 		num = num / 10
// 		count += 1
// 	}
// 	fmt.Println("the total number of digit is", count)

// }

// __________________________----
// to check whether the number is odd or even
func main() {
	var num int
	fmt.Println("ENTER ANY NUMBER TO TEST")
	fmt.Scanln(&num)
	if num%2 == 0 {
		fmt.Println("the number is even")

	} else {
		fmt.Println("the number is odd")
	}

}
