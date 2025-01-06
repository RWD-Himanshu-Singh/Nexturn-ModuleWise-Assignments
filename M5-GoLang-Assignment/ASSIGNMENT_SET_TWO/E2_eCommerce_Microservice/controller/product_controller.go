package controller

import (
	"ecommerce/model"
	"ecommerce/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProductController handles product-related HTTP requests
type ProductController struct {
	ProductService *service.ProductService
}

// NewProductController returns a new instance of ProductController
func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

// AddProduct handles POST /products requests
func (controller *ProductController) AddProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := controller.ProductService.AddProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdProduct)
}

// GetProducts handles GET /products requests
func (controller *ProductController) GetProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	product, err := controller.ProductService.GetProduct(productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct handles PUT /products/{id} requests
func (controller *ProductController) UpdateStock(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var stockUpdate struct {
		Stock int `json:"stock"`
	}
	if err := c.ShouldBindJSON(&stockUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.ProductService.UpdateStock(productID, stockUpdate.Stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock updated successfully"})
}

// DeleteProduct handles DELETE /products/{id} requests
func (controller *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = controller.ProductService.DeleteProduct(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// GetAllProducts handles GET /products requests with pagination and sorting options
func (controller *ProductController) GetAllProducts(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	products, err := controller.ProductService.GetAllProducts(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
