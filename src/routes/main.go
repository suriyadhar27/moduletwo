package main

import (
	"log"
	"net/http"
	"strconv"

	"task/src/entities"
	"task/src/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	productModel := &models.ProductModel{}

	// Route for retrieving all students
	r.GET("/students", func(c *gin.Context) {
		students, err := productModel.FindAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, students)
	})

	// Route for retrieving a specific student
	r.GET("/students/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}

		student, err := productModel.Find(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, student)
	})

	// Route for creating a new student
	r.POST("/students", func(c *gin.Context) {
		var student entities.Student
		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student data"})
			return
		}

		response := productModel.Create(&student)
		c.JSON(http.StatusOK, gin.H{"response": response})
	})

	// Route for updating a student
	r.PUT("/students/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}

		var student entities.Student
		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student data"})
			return
		}
		student.Id = id

		response, err := productModel.Update(student)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": response})
	})

	// Route for deleting a student
	r.DELETE("/students/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}

		success := productModel.Delete(id)
		c.JSON(http.StatusOK, gin.H{"success": success})
	})

	log.Println("Server started on port 8000")
	log.Fatal(r.Run(":8000"))
}
