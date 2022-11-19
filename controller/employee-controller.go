package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)

type EmployeeController interface {
	GetAll() []entity.Employee
	Add(ctx *gin.Context) entity.Employee
}

type controller struct {
	service service.EmployeeService
}

func New(service service.EmployeeService) EmployeeController {
	return &controller{
		service: service,
	}
}

func (c *controller) GetAll() []entity.Employee {
	fmt.Println("getAll employee at controller")
	return c.service.GetAll()
}

func (c *controller) Add(ctx *gin.Context) entity.Employee {
	fmt.Println("employee added at controller")
	var employee entity.Employee
	ctx.BindJSON(&employee)
	c.service.Add(employee)
	return employee
}
