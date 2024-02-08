// Q1
// You're developing an online store application in GoLang. As part of the application, you need to keep track of various product details such as name, price, and quantity in stock. Design a set of variables and assign values to represent a specific product in the inventory. Ensure you use appropriate data types for each variable to accurately capture the information.

package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func main() {
	var noOfProducts int
	fmt.Printf("Online Store Inventory Management:\n")
	fmt.Print("Enter the number of products: ")
	fmt.Scanln(&noOfProducts)

	products := make([]Product, 0, noOfProducts)

	for i := 0; i < noOfProducts; i++ {
		var (
			name     string
			price    float64
			quantity int
		)

		fmt.Print("Enter Product Name: ")
		fmt.Scan(&name)
		fmt.Print("Enter Product Price: ")
		fmt.Scan(&price)
		fmt.Print("Enter Quantity in Stock: ")
		fmt.Scan(&quantity)

		product := Product{Name: name, Price: price, Quantity: quantity}
		products = append(products, product)
	}

	fmt.Println("\nProduct details are:")
	for _, product := range products {
		fmt.Printf("{ Name: %v | Price: Rs%.2f | Quantity: %v }\n", product.Name, product.Price, product.Quantity)
	}
}
