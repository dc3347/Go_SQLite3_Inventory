package main

import(
	"fmt"
	"moodels"
)

func main(){
	var productModel models.ProductModel
	
	result := productModel.Delete(5)
	fmt.Println("result: ", result)
}