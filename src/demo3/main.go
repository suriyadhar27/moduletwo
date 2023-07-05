package main

import (
	"fmt"
	"task/src/entities"
	"task/src/models"
	"time"
)

func main() {
	var productModel models.ProductModel
	student := entities.Student{
		Id:          6,
		Title:       "agile manifesto",
		Description: "learn in depth about agile lifecycle",
		Duedate:     time.Now(),
		Status:      "in progress",
	}
	result := productModel.Create(&student)
	fmt.Println("result: ", result)
	fmt.Println(student)
}
