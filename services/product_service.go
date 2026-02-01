package services

import (
	"categories-api/model"
	"categories-api/repositories"
	"errors"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts() ([]model.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *ProductService) CreateProduct(product *model.Product) error {
	if product.CategoryID <= 0 {
		return errors.New("category_id is required")
	}

	return s.repo.CreateProduct(product)
}

func (s *ProductService) GetProductById(id int) (*model.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *ProductService) UpdateProductById(product *model.Product) error {
	if product.ID <= 0 {
		return errors.New("product ID is required for update")
	}
	if product.CategoryID <= 0 {
		return errors.New("category_id is required")
	}

	return s.repo.UpdateProductById(product)
}

func (s *ProductService) DeleteProductById(id int) error {
	if id <= 0 {
		return errors.New("invalid product ID")
	}

	return s.repo.DeleteProductById(id)
}
