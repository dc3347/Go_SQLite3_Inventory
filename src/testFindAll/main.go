package main

import(
	"fmt"
	"moodels"
)

func main(){
	var productModel models.ProductModel
	products, err := productModel.FindAll()
	if err != nil {
		fmt.Println(err)
	}
	else{
		fmt.Println("Product List")
		for _, product := range products {
			fmt.Println("id: ", product.Id)
			fmt.Println("name: ", product.Name)
			fmt.Println("price: ", product.Price)
			fmt.Println("quantity: ", product.Quantity)
			fmt.Println("supplier: ", product.Supplier)
			fmt.Println("-----------------------------")
		}
	}
}