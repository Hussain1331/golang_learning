package main

import "fmt"

func main() {
	var num int
	fmt.Println("ENTER THE NUMBER u wanna reverse:-")
	fmt.Scanln(&num)
	for i := num; i >= 1; i-- {
		fmt.Println(i, "\t")
	}

}
