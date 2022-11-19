package entity

type Employee struct {
	EmployeeName string `json : "employeeName"`
	EmployeeId   int    `json : "employeeId"`
	Designation  string `json : "designation"`
}
