package service

import (
	"fmt"

	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
)

type EmployeeService interface {
	Add(entity.Employee) entity.Employee
	GetAll() []entity.Employee
}

type employeeService struct {
	employees []entity.Employee
}

func New() EmployeeService {
	return &employeeService{}
}

func (service *employeeService) Add(employee entity.Employee) entity.Employee {
	fmt.Println("employee added in data")
	service.employees = append(service.employees, employee)

	return employee
}

func (service *employeeService) GetAll() []entity.Employee {
	fmt.Println("list of employees from data")
	return service.employees
}
