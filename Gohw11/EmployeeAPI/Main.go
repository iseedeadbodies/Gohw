package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Profession     string `json:"profession"`
	Salary         int    `json:"salary"`
	ExperienceYear int    `json:"experience_year"`
}

var employees = []Employee{
	{
		Id:             1,
		Name:           "Ali",
		Profession:     "Developer",
		Salary:         300000,
		ExperienceYear: 2,
	},
	{
		Id:             2,
		Name:           "Dias",
		Profession:     "Designer",
		Salary:         250000,
		ExperienceYear: 1,
	},
}

func main() {

	router := gin.Default()

	router.GET("/api/getallemployees", getAllEmployees)
	router.GET("/api/getone/:id", getOneEmployee)
	router.GET("/api/getnameexpyear", getNameExpYear)
	router.POST("/api/createemployee", createEmployee)

	router.Run(":8080")
}

func getAllEmployees(c *gin.Context) {
	c.JSON(http.StatusOK, employees)
}

func getOneEmployee(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	for _, employee := range employees {

		if employee.Id == id {
			c.JSON(http.StatusOK, employee)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "employee not found",
	})
}

func getNameExpYear(c *gin.Context) {

	var result []gin.H

	for _, employee := range employees {

		result = append(result, gin.H{
			"name":            employee.Name,
			"experience_year": employee.ExperienceYear,
		})
	}

	c.JSON(http.StatusOK, result)
}

func createEmployee(c *gin.Context) {

	var newEmployee Employee

	err := c.BindJSON(&newEmployee)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}

	newEmployee.Id = len(employees) + 1

	employees = append(employees, newEmployee)

	c.JSON(http.StatusCreated, newEmployee)
}
