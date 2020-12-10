package main

import(
	"fmt"
	"moodels"
)

func main(){
	var productModel models.ProductModel
	product := entities.Product{
		Name: "abc",
		Price: 11.5,
		Quantity: 9,
		Status: true,
	}
	result := productModel.Create(&product)
	fmt.Println("result: ", result)
	fmt.Println(product)
}