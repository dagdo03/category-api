package repository

import (
	"categories-api/model"
)

var categories = []model.Category{
	{ID: 1, Name: "Electronics", Description: "Devices and gadgets"},
	{ID: 2, Name: "Books", Description: "Printed and digital books"},
	{ID: 3, Name: "Clothing", Description: "Apparel and accessories"},
	{ID: 4, Name: "Home & Kitchen", Description: "Household items and kitchenware"},
	{ID: 5, Name: "Sports", Description: "Sporting goods and outdoor equipment"},
	{ID: 6, Name: "Toys", Description: "Toys and games for children"},
	{ID: 7, Name: "Beauty", Description: "Cosmetics and personal care products"},
	{ID: 8, Name: "Health", Description: "Healthcare and wellness products"},
	{ID: 9, Name: "Jewelry", Description: "Jewelry and accessories"},
	{ID: 10, Name: "Automotive", Description: "Automotive parts and accessories"},
}

func GetCategories() []model.Category {
	return categories
}

func AddCategory(category model.Category) model.Category {
	category.ID = len(categories) + 1
	categories = append(categories, category)
	return category
}

func DeleteCategoryById(id int) bool {
	for index, category := range categories {
		if category.ID == id {
			categories = append(categories[:index], categories[index+1:]...)
			return true
		}
	}
	return false
}

func UpdateCategoryById(id int, category model.Category) bool {
	for index, cat := range categories {
		if cat.ID == id {
			category.ID = id
			categories[index] = category
			return true
		}
	}
	return false
}

func GetCategoryById(id int) *model.Category {
	for _, category := range categories {
		if category.ID == id {
			return &category
		}
	}
	return nil
}
