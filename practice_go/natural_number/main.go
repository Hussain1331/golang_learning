package main

import "fmt"

func main() {
	var num int
	fmt.Println("ENTER THE NUMBER LIMIT :-")
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		fmt.Println(i, "\t")
	}

}
