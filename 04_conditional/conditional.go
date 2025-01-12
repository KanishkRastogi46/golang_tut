package main

import (
	"fmt"
	"time"
)

func main() {
	// if..elseif..else statement
	var age int
	fmt.Printf("Enter your age: ")
	fmt.Scanf("%d", &age)
	fmt.Println("if..elseif..else statement")
	if age < 18 {
		fmt.Println("You are a minor, your age is", age)
	} else if age >= 100 || age < 0 {
		fmt.Println("Invalid age")
	} else {
		fmt.Println("You are an adult, your age is", age)
	}

	// switch-case statement
	var num int
	fmt.Printf("Enter a number from 1-3: ")
	fmt.Scan(&num)
	fmt.Println("switch-case statement")
	switch num {
		case 1:
			fmt.Println("Your value is 1")
		case 2:
			fmt.Println("Your value is 2")
		case 3:
			fmt.Println("Your value is 3")
		default:
			fmt.Println("Invalid value")
	}

	// multiple conditions
	fmt.Println("Multiple switch conditions")
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("It's the weekend")
		default:
			fmt.Println("It's a weekday")
	}
}