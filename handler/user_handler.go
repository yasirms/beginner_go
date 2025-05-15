package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/yasirms/beginner_go/model"
	"github.com/yasirms/beginner_go/repository"

	"github.com/gin-gonic/gin"
)

func GetAllEmployees(c *gin.Context) {
	employees, err := repository.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees"})
		return
	}
	c.JSON(http.StatusOK, employees)
}

func CreateEmployee(c *gin.Context) {
	var emp model.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := repository.CreateEmployee(emp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	// id, _ = strconv.Itoa(id)
	fmt.Println(id)
	employee, err := repository.GetEmployeeByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employee"})
		return
	}
	// fmt.Println(employee)
	c.JSON(http.StatusOK, employee)
}

func UpdateEmployee(c *gin.Context) {
	var emp model.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	err := repository.UpdateEmployee(emp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employee"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully"})
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	id_string, _ := strconv.Atoi(id)
	err := repository.DeleteEmployee(id_string)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}

func HandleLogin(c *gin.Context) {
	var cred model.UserCredentials
	if err := c.ShouldBindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	emp, err := repository.ValidateLogin(cred.Email, cred.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}
	c.JSON(http.StatusOK, emp)
}