package services

import (
	"categories-api/model"
	"categories-api/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAllCategories() ([]model.Category, error) {
	return s.repo.GetAllCategories()
}

func (s *CategoryService) CreateCategory(data *model.Category) error {
	return s.repo.CreateCategory(data)
}

func (s *CategoryService) UpdateCategoryById(category *model.Category) error {
	return s.repo.UpdateCategoryById(category)
}

func (s *CategoryService) GetCategoryById(id int) (*model.Category, error) {
	return s.repo.GetCategoryByID(id)
}

func (s *CategoryService) DeleteCategoryById(id int) error {
	return s.repo.DeleteCategoryById(id)
}
