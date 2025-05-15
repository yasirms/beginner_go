package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yasirms/beginner_go/config"
	"github.com/yasirms/beginner_go/handler"
)

func main() {

	config.InitDB()
	r := gin.Default()
	r.GET("/employees", handler.GetAllEmployees)
	r.GET("/employees/:id", handler.GetEmployeeByID)
	r.POST("/employees", handler.CreateEmployee)
	r.PUT("/employees/:id", handler.UpdateEmployee)
	r.DELETE("/employees/:id", handler.DeleteEmployee)
	r.Run("localhost:8090")

}
