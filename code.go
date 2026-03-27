package main

import (
	"fmt"
)

func main() {
	var number1, number2 int
	//var operation string
	var operator string

	for {
		fmt.Println("Welcome to my Calculator Program")
		fmt.Println("Insert the number for operation")
		fmt.Print("Enter first number: ")
		fmt.Scanln(&number1)
		fmt.Println()
		fmt.Println("Enter operator: addition, substraction, multiplication, division ")
		fmt.Scanln(&operator)
		fmt.Print("Enter second number: ")
		fmt.Scanln(&number2)
		fmt.Print()

		switch operator {
		case "+":
			fmt.Println("Final Result = ", number1+number2)

		case "-":
			fmt.Println("Final Result =", number1-number2)

		case "*":
			fmt.Println("Final Result =", number1*number2)

		case "/":
			if number2 == 0 {
				fmt.Println("Leave here cannot divide by 0")
			} else {
				fmt.Println("Final Result =", number1/number2)
			}
		default:
			fmt.Println("Leave Here: invalid input")

		}

		continue

	}

}
