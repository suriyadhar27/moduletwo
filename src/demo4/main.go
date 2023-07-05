package main

import (
	"fmt"
	"task/src/models"
)

func main() {
	var productModel models.ProductModel
	result := productModel.Delete(6)
	fmt.Println("result: ", result)
}
