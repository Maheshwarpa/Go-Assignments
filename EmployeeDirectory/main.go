package main

import (
	"github.com/Maheshwarpa/Go-Assignments/repository"
	"github.com/Maheshwarpa/Go-Assignments/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

/*
CRUD

*/

func main() {

	repo := repository.NewEmployeeRepo()

	Execute(repo)

}

func Execute(repo service.EmployeeService) {

	//Create

	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodPost:
			{

				repo.CreateEmployee(w, r)

			}
		case http.MethodPut:
			{

				repo.UpdateEmployee(w, r)

			}
		case http.MethodGet:
			{
				// var x int
				// fmt.Println("Please enter the value of X")
				// fmt.Scan(&x)
				// fmt.Println(repo.GetEmployee(x))
				idStr := r.URL.Query().Get("id") // Extract 'id' from query parameters
				id, err := strconv.Atoi(idStr)   // Convert to integer
				if err != nil {
					http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
					return
				}
				employee, err := repo.GetEmployee(id)
				if err != nil {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				json.NewEncoder(w).Encode(employee)

			}
		case http.MethodDelete:
			{
				// var x int
				// fmt.Println("Please enter the value of X")
				// fmt.Scan(&x)
				// fmt.Println(repo.GetEmployee(x))
				idStr := r.URL.Query().Get("id") // Extract 'id' from query parameters
				id, err := strconv.Atoi(idStr)   // Convert to integer
				if err != nil {
					http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
					return
				}
				//Before
				fmt.Println("******************************")
				//repo.ListAllEmployees()
				//While Delete
				repo.DeleteEmployee(id)
				//After Delete
				//repo.ListAllEmployees()

			}
		default:
			{
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}

		}
	})

	fmt.Println("Starting Server")

	http.ListenAndServe(":8080", nil)
	/*
		//Update

		err := repo.UpdateEmployee(models.Employee{
			EmployeeID:   2,
			EmployeeName: "Bhavani",
			EmployeeAge:  30,
		})

		if err != nil {
			repo.ListAllEmployees()
			fmt.Println(err)
		} else {
			fmt.Println("Employee Updated Successfully ")
		}

		//GET
		fmt.Println("**********Getting an Employee with ID 2***************")
		val, err := repo.GetEmployee(2)

		fmt.Println(val)

		if err != nil {
			repo.ListAllEmployees()
			fmt.Println(err)
		} else {
			fmt.Println("Employee Updated Successfully ")
		}

		//Delete

		fmt.Println("**********Deleting an Employee with ID 2***************")
		err = repo.DeleteEmployee(2)

		if err != nil {
			repo.ListAllEmployees()
			fmt.Println(err)
		} else {
			fmt.Println("Employee Updated Successfully ")
		}
	*/

}
