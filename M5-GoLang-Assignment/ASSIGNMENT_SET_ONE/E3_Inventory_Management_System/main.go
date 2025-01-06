package main

import (
	"Inventory/services"
	"fmt"
)

func options() {
	fmt.Println("\n=== Inventory Management System ===")
	fmt.Println("1) Add Product")
	fmt.Println("2) Update Stock")
	fmt.Println("3) Search Product")
	fmt.Println("4) Display Inventory")
	fmt.Println("5) Sort by Price")
	fmt.Println("6) Sort by Stock")
	fmt.Println("7) Exit")
	fmt.Print("Enter your choice: ")
}

func main() {
	for {
		options()
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Add Product
			var name, price string
			var stock int
			fmt.Print("Enter product name: ")
			fmt.Scan(&name)
			fmt.Print("Enter product price: ")
			fmt.Scan(&price)
			fmt.Print("Enter product stock: ")
			fmt.Scan(&stock)

			err := services.AddProduct(name, price, stock)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Product added successfully.")
			}

		case 2:
			// Update Stock
			var productID, newStock int
			fmt.Print("Enter product ID: ")
			fmt.Scan(&productID)
			fmt.Print("Enter new stock: ")
			fmt.Scan(&newStock)

			err := services.UpdateStock(productID, newStock)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Stock updated successfully.")
			}

		case 3:
			// Search Product
			var identifier string
			fmt.Print("Enter product ID or name: ")
			fmt.Scan(&identifier)

			product, err := services.SearchProduct(identifier)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Product Found: %+v\n", *product)
			}

		case 4:
			// Display Inventory
			fmt.Println("\nCurrent Inventory:")
			services.DisplayInventory()

		case 5:
			// Sort by Price
			ans := services.SortByPrice()
			fmt.Println("\nProducts sorted by price:")
			fmt.Println("ID\tName\t\tPrice\t\tStock")
			fmt.Println("----------------------------------------")
			for _, product := range ans {
				fmt.Printf("%d\t%s\t\t%.2f\t\t%d\n", product.ID, product.Name, product.Price, product.Stock)
			}

		case 6:
			// Sort by Stock
			ans := services.SortByStock()
			fmt.Println("\nProducts sorted by stocks:")
			fmt.Println("ID\tName\t\tPrice\t\tStock")
			fmt.Println("----------------------------------------")
			for _, product := range ans {
				fmt.Printf("%d\t%s\t\t%.2f\t\t%d\n", product.ID, product.Name, product.Price, product.Stock)
			}

		case 7:
			// Exit
			fmt.Println("Exiting Inventory Management System. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
