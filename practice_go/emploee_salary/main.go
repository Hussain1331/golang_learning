package main

import "fmt"

func main() {
	var basic_salary float32
	fmt.Println("ENTER YOUR SALARY")
	fmt.Scanln(&basic_salary)
	hra := basic_salary * 0.20
	da := basic_salary * 0.10
	gross := hra + da + basic_salary
	println("YOOUR BASIC SALARY", basic_salary)
	println("HOUSE RENT ALLOWANCE", hra)
	println("dearness allowance", da)
	println("so your gross salary", gross)

}
