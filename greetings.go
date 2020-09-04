package main

import "fmt"

func main() {

	//fmt.Println("Soma de 5 + 5")
	fmt.Print("Enter first number: ")
	var input1 int
	fmt.Scanln(&input1)
	fmt.Print("Enter secont number: ")
	var input2 int
	fmt.Scanln(&input2)
	fmt.Println("resultado da soma ", input1, " + ", input2, " = ", sum(input1, input2))
}
