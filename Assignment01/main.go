package main

import (
	"fmt"
	"github.com/Maheshwarpa/Go-Assignments/structPractice"
)

func main() {
	//Hello World!
	var s structPractice.Student
	//Hi Hello
	s.Name = "Maheshwar"
	s.Marks = []int{10, 10, 10, 10, 10}
	var king int
	fmt.Scan(&king)
	fmt.Println("The King Value is :",king)
	fmt.Println("Average is:", s.CalculateAverage())
	s.AddMark(20)
	fmt.Println(s.Name)
	fmt.Println(s.Marks)
	fmt.Printf("Average is: %.2f", s.CalculateAverage())

}
