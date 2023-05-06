package main

import "fmt"

func main() {
	num1 := 20045
	num2 := 60000000
	num3 := 1
	sum := num1 + num2 + num3
	fmt.Println("The sum of", +num1, +num2, "and", +num3, "is", sum)

	//the heighest and lowest numbers
	max := num1
	if num2 > max {
		max = num2
	}
	if num3 > max {
		max = num3
	}

	//lowest numbers
	min := num1
	if num2 < min {
		min = num2
	}
	if num3 < min {
		min = num3
	}
	fmt.Println("The heighest number is", max, "and", "The smallest number is", min)
}
