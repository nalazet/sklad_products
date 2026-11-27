package storage

import (
	"errors"
	"sklad_products/models"
)

func AddProduct(products []models.Product, product models.Product) ([]models.Product, error) {
	for _, p := range products {
		if p.ID == product.ID {
			return products, errors.New("Продукт с таким ID уже существует")
		}
	}

	products = append(products, product)

	return products, nil
}

func FindProductByID(products []models.Product, id models.ProductID) (models.Product, error) {
	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}

	return models.Product{}, errors.New("Продукт не найден")
}
