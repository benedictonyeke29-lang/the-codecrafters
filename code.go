package main

import (
	"fmt"
)

func main() {
	var number1, number2 float64
	//var operation string
	var operator string

	for {
		fmt.Println("Welcome to my Calculator Program")
		fmt.Println("Insert the number for operation")
		fmt.Print("Enter first number: ")
		fmt.Scanln(&number1)
		fmt.Println()
		fmt.Print("Enter second number: ")
		fmt.Scanln(&number2)
		fmt.Print()
		fmt.Println("Enter operator: +, -, *, / or type quit to exit ")
		fmt.Scanln(&operator)

       if operator == "help" {
		fmt.Println("Expected command: +, -, *, / " )
	   }

	   if operator == "quit" {
		fmt.Println("Good bye!!")
		break
	   }


		switch operator {
		case "+":
			fmt.Println("Final Result = ", number1+number2)
			continue

		case "-":
			fmt.Println("Final Result =", number1-number2)
			continue

		case "*":
			fmt.Println("Final Result =", number1*number2)
			continue

		case "/":
			if number2 == 0 {
				fmt.Println("Leave here cannot divide by 0")
				continue

			} else {
				fmt.Println("Final Result =", number1/number2)
			}
		default:
			fmt.Println("Leave Here: invalid input")

		}

	}

}
