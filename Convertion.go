package main

import "fmt"

func main() {
	convert := 20.6
	celsius := celsius(convert)
	fmt.Printf("%.2f°F is %.2f°C", convert, celsius)
}
func celsius(convert float64) float64 {
	celsius := (convert - 32) * 5 / 9
	return celsius

}
