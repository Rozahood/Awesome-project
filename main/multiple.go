package main

import "fmt"

func main() {

	limit := 10

	for i := 1; i <= limit; i++ {

		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
}
