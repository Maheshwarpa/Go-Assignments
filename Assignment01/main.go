package main

import (
	"fmt"
	"module/structPractice"
)

func main() {

	var s structPractice.Student

	s.Name = "Maheshwar"
	s.Marks = []int{10, 10, 10, 10, 10}

	fmt.Println("Average is:", s.CalculateAverage())
	s.AddMark(20)
	fmt.Println(s.Name)
	fmt.Println(s.Marks)
	fmt.Println("Average is:", s.CalculateAverage())

}