package services

import (
	"Inventory/models"
	"errors"
	"fmt"
	"sort"
	"strconv"
)

var inventory []models.Product

// Adding Product
func AddProduct(name string, priceInput string, stock int) error {
	price, err := strconv.ParseFloat(priceInput, 64)
	if err != nil {
		return errors.New("invalid price format")
	}
	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	id := len(inventory) + 1
	inventory = append(inventory, models.Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	})
	return nil
}

// Update Stock
func UpdateStock(productID int, newStock int) error {
	if newStock < 0 {
		return errors.New("stock can't be negative")
	}

	for i := range inventory {
		if inventory[i].ID == productID {
			inventory[i].Stock = newStock
			return nil
		}
	}
	return errors.New("product not found")
}

// Search Product by ID or Name
func SearchProduct(identifier string) (*models.Product, error) {
	id, err := strconv.Atoi(identifier)
	if err == nil {
		// Search by ID
		for i := range inventory {
			if inventory[i].ID == id {
				return &inventory[i], nil
			}
		}
	} else {
		// Search by Name
		for i := range inventory {
			if inventory[i].Name == identifier {
				return &inventory[i], nil
			}
		}
	}
	return nil, errors.New("product not found")
}

// returns a list of products
func DisplayInventory() {
	fmt.Println("ID\tName\t\tPrice\t\tStock")
	fmt.Println("----------------------------------------")
	for _, product := range inventory {
		fmt.Printf("%d\t%s\t\t%.2f\t\t%d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

// Sort by Price
func SortByPrice() []models.Product {
	tempProducts := make([]models.Product, len(inventory))
	copy(tempProducts, inventory)

	// Sort the copied slice
	sort.Slice(tempProducts, func(i, j int) bool {
		return tempProducts[i].Price < tempProducts[j].Price
	})

	return tempProducts
}

// Sort by Stock
func SortByStock() []models.Product {
	// Create a copy of the original slice
	tempProducts := make([]models.Product, len(inventory))
	copy(tempProducts, inventory)

	// Sort the copied slice
	sort.Slice(tempProducts, func(i, j int) bool {
		return tempProducts[i].Stock < tempProducts[j].Stock
	})

	return tempProducts
}
