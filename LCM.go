package main

import "fmt"

func main() {
	num := []int{98, 78, 456, 5, 12}

	lcm := calculateLCM(num)
	fmt.Printf("LCM of %v is %d/n", num, lcm)
}

func calculateLCM(num []int) int {
	lcm := num[0]
	for i := 1; i < len(num); i++ {
		lcm = calculateTwoNumLCM(lcm, num[i])
	}
	return lcm
}

func calculateTwoNumLCM(num1, num2 int) int {
	lcm := calculateTwoNumLCM(num1, num2)
	return (num1 * num2) / lcm
}
