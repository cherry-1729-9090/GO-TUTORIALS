package main

import "fmt"

func DecisionMaking() {
	num := 15
	if num%3 == 0 {
		fmt.Println("Number is multiple of 3")
	} else if num%5 == 0 {
		fmt.Println("Number is multiple of 5")
	} else {
		fmt.Println("Number is not multiple of 3 or 5")
	}

	var age int
	fmt.Scanln(&age)
	switch age {
	case 10:
		fmt.Println("You are 10 years old")

	case 20:
		fmt.Println("You are 20 years old")

	case 30:
		fmt.Println("You are 30 years old")

	default:
		fmt.Println("You are not 10, 20 or 30 years old")
	}

}
