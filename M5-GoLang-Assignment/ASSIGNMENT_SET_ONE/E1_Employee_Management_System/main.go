package main

import (
	"Ex1_Emp_Man/services"
	"Ex1_Emp_Man/utils"
	"fmt"
)

func main() {
	// Adding sample employees
	services.AddEmployee("John Doe", 25, utils.IT)
	services.AddEmployee("Jane Smith", 30, utils.HR)
	services.AddEmployee("Alice Brown", 22, utils.IT)

	// Count employees in IT department
	itCount := services.CountEmployees(utils.IT)
	fmt.Printf("Number of employees in IT: %d\n", itCount)

	// Count employees in HR department
	hrCount := services.CountEmployees(utils.HR)
	fmt.Printf("Number of employees in HR: %d\n", hrCount)

	// Attempt to count employees in an invalid department
	otherCount := services.CountEmployees("Operations")
	fmt.Printf("Number of employees in Operations: %d\n", otherCount)
}
