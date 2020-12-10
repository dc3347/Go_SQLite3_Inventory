package main

import(
	"fmt"
	"moodels"
)

func main(){
	var productModel models.ProductModel
	product, _ := productModel.Find(3)
	product.Name = "xyz"
	product.Price = 123
	product.Supplier = "q"
	result := productModel.Update(product)
	fmt.Println("result: ", result)
}