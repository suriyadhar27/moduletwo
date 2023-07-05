package main

import (
	"fmt"
	"log"
	"task/src/config"
	"task/src/models"
)

func main() {
	var productModel models.ProductModel
	config.GetDB()
	students, err := productModel.FindAll()
	log.Println(err)
	log.Println(len(students))
	if err != nil {
		log.Println(err)
		return
	} else {
		fmt.Println("Student list:")
		for _, student := range students {
			fmt.Println("Id:", student.Id)
			fmt.Println("Title:", student.Title)
			fmt.Println("Description:", student.Description)
			fmt.Println("Duedate:", student.Duedate)
			fmt.Println("Status:", student.Status)
			fmt.Println("-----------------------------------------")
		}
	}
}
