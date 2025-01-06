package service

import (
	"ecommerce/model"
	"ecommerce/repository"
)

type ProductService struct {
	ProductRepo *repository.ProductRepository
}

func NewProductService(productRepo *repository.ProductRepository) *ProductService {
	return &ProductService{ProductRepo: productRepo}
}

func (service *ProductService) AddProduct(product *model.Product) (*model.Product, error) {
	return service.ProductRepo.AddProduct(product)
}

func (service *ProductService) GetProduct(id int) (*model.Product, error) {
	return service.ProductRepo.GetProduct(id)
}

func (service *ProductService) UpdateStock(id int, stock int) error {
	return service.ProductRepo.UpdateStock(id, stock)
}

func (service *ProductService) DeleteProduct(id int) error {
	return service.ProductRepo.DeleteProduct(id)
}

func (service *ProductService) GetAllProducts(page, limit int) ([]model.Product, error) {
	return service.ProductRepo.GetAllProducts(page, limit)
}
