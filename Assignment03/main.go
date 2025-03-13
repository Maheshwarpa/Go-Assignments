package main

import (
	"fmt"
	"github.com/Maheshwarpa/Go-Assignments.git/Employee"
)

func main() {
	var Emp Employee.EmployeeList
	emp := Employee.Employee{1, "Maheshwar", "Chennai", 75000}
	emp1 := Employee.Employee{2, "Maheshwar", "Chennai", 75000}
	emp2 := Employee.Employee{3, "Maheshwar", "Chennai", 75000}
	emp3 := Employee.Employee{4, "Maheshwar", "Chennai", 75000}
	emp4 := Employee.Employee{5, "Maheshwar", "Chennai", 75000}
	emp5 := Employee.Employee{6, "Maheshwar", "Chennai", 75000}
	emp6 := Employee.Employee{7, "Maheshwar", "Chennai", 75000}

	Emp.AddEmployee(emp)
	Emp.AddEmployee(emp1)
	Emp.AddEmployee(emp2)
	Emp.AddEmployee(emp2)
	Emp.AddEmployee(emp3)

	fmt.Println("*****************")
	Emp.DisplayAll()
	fmt.Println("*****************")
	Emp.DeleteEmployee(3)
	fmt.Println("*****************")
	Emp.DisplayAll()
	Emp.AddEmployee(emp4)
	Emp.AddEmployee(emp5)
	Emp.UpdateEmployee(5, emp6)
	fmt.Println("*****************")
	fmt.Println(Emp.DisplayEmployee(7))
	fmt.Println("*****************")
	Emp.DisplayAll()

}
