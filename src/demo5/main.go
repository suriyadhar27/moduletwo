package main

import (
	"fmt"
	"task/src/models"
	"time"
)

func main() {
	var productModel models.ProductModel
	student, _ := productModel.Find(3)
	student.Title = "def"
	student.Description = "to know more about it"
	student.Duedate = time.Now()
	student.Status = "in progress"
	result, err := productModel.Update(student)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

// package main

// import (
// 	"fmt"
// 	"task/src/models"
// 	"time"
// )

// func main() {
// 	var productModel models.ProductModel
// 	student, _ := productModel.Find(3)
// 	student.Title = "def"
// 	student.Description = "to know more about it"
// 	student.Duedate = time.Now()
// 	student.Status = "in progress"
// 	result := productModel.Update(student)
// 	fmt.Println("result:", result)
// }
