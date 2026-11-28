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

func FindCategoryByID(categories []models.Category, id models.CategoryID) (models.Category, error) {
	for _, category := range categories {
		if category.ID == id {
			return category, nil
		}
	}

	return models.Category{}, errors.New("Категория не найдена")
}

func UpdateCategory(categories []models.Category, updatedCategory models.Category) error {
	for i, category := range categories {
		if category.ID == updatedCategory.ID {
			categories[i] = updatedCategory
			return nil
		}
	}

	return errors.New("Категория для обновления не найдена")
}
