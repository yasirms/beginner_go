package repository

import (
	"github.com/yasirms/beginner_go/config"
	"github.com/yasirms/beginner_go/model"
)

func GetAllEmployees() ([]model.Employee, error) {
	rows, err := config.DB.Query("SELECT * FROM employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []model.Employee
	for rows.Next() {
		var emp model.Employee
		if err := rows.Scan(&emp.ID, &emp.EmployeeID, &emp.Name, &emp.FatherName, &emp.Email, &emp.Address, &emp.PhoneNumber, &emp.CNICNumber, &emp.JobTitle, &emp.JobStartDate, &emp.JobEndDate, &emp.CreatedAt, &emp.UpdatedAt); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, nil
}

func CreateEmployee(emp model.Employee) (int, error) {
	var id int
	err := config.DB.QueryRow("INSERT INTO employee (employee_id, name, father_name, email, address, phone_number, cnic_number, job_title, job_start_date, job_end_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id",
		emp.EmployeeID,
		emp.Name,
		emp.FatherName,
		emp.Email,
		emp.Address,
		emp.PhoneNumber,
		emp.CNICNumber,
		emp.JobTitle,
		emp.JobStartDate,
		emp.JobEndDate).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func UpdateEmployee(emp model.Employee) error {
	_, err := config.DB.Exec("UPDATE employee SET employee_id = $1, name = $2, father_name = $3, email = $4, address = $5, phone_number = $6, cnic_number = $7, job_title = $8, job_start_date = $9, job_end_date = $10 WHERE id = $11",
		emp.EmployeeID,
		emp.Name,
		emp.FatherName,
		emp.Email,
		emp.Address,
		emp.PhoneNumber,
		emp.CNICNumber,
		emp.JobTitle,
		emp.JobStartDate,
		emp.JobEndDate,
		emp.ID)
	return err
}

func DeleteEmployee(id int) error {
	_, err := config.DB.Exec("DELETE FROM employee WHERE id = $1", id)
	return err
}

func GetEmployeeByID(id int) (model.Employee, error) {
	var emp model.Employee
	err := config.DB.QueryRow("SELECT * FROM employee WHERE id = $1", id).Scan(&emp.ID, &emp.EmployeeID, &emp.Name, &emp.FatherName, &emp.Email, &emp.Address, &emp.PhoneNumber, &emp.CNICNumber, &emp.JobTitle, &emp.JobStartDate, &emp.JobEndDate, &emp.CreatedAt, &emp.UpdatedAt)
	if err != nil {
		return emp, err
	}
	return emp, nil
}
