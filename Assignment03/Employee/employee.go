package Employee

import "fmt"

type Employee struct {
	Emp_Id   int
	Emp_Name string
	Emp_Loc  string
	Emp_Sal  int
}

type EmployeeList []Employee

func (e *EmployeeList) AddEmployee(ad Employee) []Employee {
	var flag bool
	flag = true
	for _, v := range *e {
		if v.Emp_Id == ad.Emp_Id {
			flag = false
		}
	}
	if flag {
		*e = append(*e, ad)
		return *e
	}
	return *e
}

func (e *EmployeeList) UpdateEmployee(id int, up Employee) []Employee {
	var flag bool
	for ind, v := range *e {
		if v.Emp_Id == id {
			(*e)[ind] = up
			flag = false
		}
	}
	if flag {
		*e = append(*e, up)
		return *e
	}
	return *e
}

func (e *EmployeeList) DeleteEmployee(id int) bool {
	var flag bool
	//var k EmployeeList
	for ind, v := range *e {
		if v.Emp_Id == id {
			flag = true
			*e = append((*e)[:ind], (*e)[(ind+1):]...)
		}
	}

	return flag

}

func (e *EmployeeList) DisplayEmployee(id int) Employee {
	var ans Employee
	for _, v := range *e {
		if v.Emp_Id == id {
			return v
		}
	}
	return ans
}

func (e *EmployeeList) DisplayAll() {
	for ind, v := range *e {
		fmt.Println(ind+1, " ", v)
	}
}
