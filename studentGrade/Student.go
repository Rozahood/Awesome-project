package main

import "fmt"

func main() {
	type Student struct {
		Name  string
		Grade float64
	}
	students := []Student{}

	for i := 0; i < len(students); i++ {
		fmt.Printf("Enter your name amd grade for student %d:/n\n", i+1)
		var name string
		var grade float64
		fmt.Scan(&name, &grade)
		students = append(students, Student{Name: name, Grade: grade})
	}

	for _, students := range students {
		fmt.Printf("%s: .2f/n", students.Name, students.Grade)
	}
}
