package model

type Employee struct {
	ID           int    `json:"id"`
	EmployeeID   string `json:"employee_id"`
	Name         string `json:"name"`
	FatherName   string `json:"father_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Role		 string `json:"role"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phone_number"`
	CNICNumber   string `json:"cnic_number"`
	JobTitle     string `json:"job_title"`
	JobStartDate string `json:"job_start_date"`
	JobEndDate   string `json:"job_end_date"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type EmployeeSalary struct {
	ID               int    `json:"id"`
	EmployeeID       string `json:"employee_id"`
	Salary           int32  `json:"salary"`
	TravelAllowance  int32  `json:"travel_allowance"`
	MedicalAllowance int32  `json:"medical_allowance"`
	AnnualLeave      int32  `json:"annual_leave"`
	SickLeave        int32  `json:"sick_leave"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type EmployeeLeave struct {
	ID          int    `json:"id"`
	EmployeeID  string `json:"employee_id"`
	LeaveType   string `json:"leave_type"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	LeaveStatus string `json:"leave_status"`
	LeaveReason string `json:"leave_reason"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UserCredentials struct {
	Email    string `json:"email" binding:"required, email"`
	Password string `json:"password" binding:"required"`
}