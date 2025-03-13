package main

import (
	"fmt"
	"github.com/Maheshwarpa/Go-Assignments/structPractice"
)

func main() {

	var s structPractice.Student

	s.Name = "Maheshwar"
	s.Marks = []int{10, 10, 10, 10, 10

	fmt.Println("Average is:", s.CalculateAverage())
	s.AddMark(20)
	fmt.Println(s.Name)
	fmt.Println(s.Marks)
	fmt.Printf("Average is: %.2f", s.CalculateAverage())

}
