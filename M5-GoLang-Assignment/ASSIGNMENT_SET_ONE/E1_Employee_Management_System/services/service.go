package services

import (
	"Ex1_Emp_Man/model"
	"errors"
	"strconv"
	"strings"
)

var employees []model.Employee

func AddEmployee(name string, age int, department string) {
	id := len(employees) + 1
	employees = append(employees, model.Employee{
		ID:         id,
		Name:       name,
		Age:        age,
		Department: department,
	})
}

func GetEmployee(identifier string) (model.Employee, error) {
	id, err := strconv.Atoi(identifier)
	if err == nil {
		for _, emp := range employees {
			if emp.ID == id {
				return emp, nil
			}
		}
	} else {
		for _, emp := range employees {
			if strings.ToLower(emp.Name) == strings.ToLower(identifier) {
				return emp, nil
			}
		}
	}
	return model.Employee{}, errors.New("employee not found")
}

func ListEmployeesByDepartment(department string) []model.Employee {
	var deptEmployees []model.Employee
	for _, emp := range employees {
		if emp.Department == department {
			deptEmployees = append(deptEmployees, emp)
		}
	}
	return deptEmployees
}

func CountEmployees(department string) int {
	count := 0
	for _, emp := range employees {
		if emp.Department == department {
			count++
		}
	}
	return count
}
