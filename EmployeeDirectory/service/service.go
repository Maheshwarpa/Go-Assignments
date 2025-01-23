package service

import (
	"employeeeDirectory/models"
	"net/http"
)

type EmployeeService interface {
	CreateEmployee(w http.ResponseWriter, r *http.Request)         //Create
	GetEmployee(id int) (models.Employee, error)                   //Read
	UpdateEmployee(w http.ResponseWriter, req *http.Request) error //Update
	DeleteEmployee(id int) error                                   //Delete
	ListAllEmployees()
}
