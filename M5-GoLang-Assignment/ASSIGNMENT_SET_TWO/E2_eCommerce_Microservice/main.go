package main

import (
	"ecommerce/config"
	"ecommerce/controller"
	"ecommerce/middleware"
	"ecommerce/repository"
	"ecommerce/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	config.InitializeDatabase()

	// Create the repository, service, and controller for the Product
	productRepo := repository.NewProductRepository(config.GetDB())
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	// Initialize Gin router
	r := gin.Default()

	// Applied the LoggingMiddleware globally
	r.Use(middleware.LoggingMiddleware())

	// Globally Applied Basic Authentication middleware to the routes
	r.Use(middleware.BasicAuthMiddleware())

	// Product routes
	r.POST("/product", productController.AddProduct)
	r.GET("/product/:id", productController.GetProduct)
	r.PUT("/product/:id", productController.UpdateStock)
	r.DELETE("/product/:id", productController.DeleteProduct)
	r.GET("/products", productController.GetAllProducts)

	// Start the server
	r.Run(":8080")
}
