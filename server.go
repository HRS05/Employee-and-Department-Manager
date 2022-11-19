package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/controller"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)

var (
	employeeService    service.EmployeeService       = service.New()
	employeeController controller.EmployeeController = controller.New(employeeService)
)

func main() {
	server := gin.Default()

	server.GET("/getEmployees", func(ctx *gin.Context) {
		fmt.Println("employee added request")
		ctx.JSON(200, employeeController.GetAll())
	})

	server.POST("/addEmployee", func(ctx *gin.Context) {
		fmt.Println("getAll employee request")
		ctx.JSON(200, employeeController.Add(ctx))
	})

	server.Run(":8080")

}
