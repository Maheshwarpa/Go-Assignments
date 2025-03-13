package repository

import (
	"github.com/Maheshwarpa/Go-Assignments.git/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type EmployeeRepo struct {
	employees map[int]models.Employee
}

func NewEmployeeRepo() *EmployeeRepo {
	return &EmployeeRepo{employees: make(map[int]models.Employee)}
}

func (r *EmployeeRepo) CreateEmployee(w http.ResponseWriter, req *http.Request) {

	var emp models.Employee

	if err := json.NewDecoder(req.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}

	fmt.Println(emp)

	if _, exists := r.employees[emp.ID()]; exists {
		fmt.Println("Employee Already Exists")
		http.Error(w, "Invalid Request body", http.StatusUnprocessableEntity)
		return
	}

	r.employees[emp.ID()] = emp

	r.ListAllEmployees()

	w.WriteHeader(http.StatusCreated)

}

func (r *EmployeeRepo) GetEmployee(id int) (models.Employee, error) {

	if val, exists := r.employees[id]; !exists {
		fmt.Println("No Employee Found")
		return models.Employee{}, errors.New("Invalid Employee ID")
	} else {
		for _, v := range *&r.employees {
			if v.EmployeeID == id {
				return v, nil
			}
		}
		return val, nil
	}

}

func (r *EmployeeRepo) UpdateEmployee(w http.ResponseWriter, req *http.Request) error {

	var emp models.Employee

	if err := json.NewDecoder(req.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)

	}

	if _, exists := r.employees[emp.ID()]; !exists {
		fmt.Println("No Employee Found to update")
		return errors.New("Invalid Employee ID")
	}

	r.employees[emp.ID()] = emp

	r.ListAllEmployees()

	w.WriteHeader(http.StatusAccepted)
	return nil
}

func (r *EmployeeRepo) DeleteEmployee(id int) error {

	if _, exists := r.employees[id]; !exists {
		fmt.Println("No Employee found to delete")
		return errors.New("Not a new Employee")
	}

	delete(r.employees, id)

	r.ListAllEmployees()
	return nil
}

func (r *EmployeeRepo) ListAllEmployees() {
	fmt.Println(r.employees)
}
