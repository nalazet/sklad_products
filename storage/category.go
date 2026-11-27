package storage

import (
	"errors"
	"sklad_products/models"
)

func AddCategory(categories []models.Category, category models.Category) ([]models.Category, error) {
	for _, c := range categories {
		if c.ID == category.ID {
			return categories, errors.New("Категория с таким ID уже существует")
		}
	}

	categories = append(categories, category)

	return categories, nil
}
