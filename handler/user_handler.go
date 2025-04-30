package handler

import (
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	id, err := repository.CreateEmployee(emp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func UpdateEmployee(c *gin.Context) {
	var emp model.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// id, _ := strconv.Atoi(c.Param("id"))
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
