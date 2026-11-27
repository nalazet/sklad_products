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

func UpdateProduct(products []models.Product, updatedProduct models.Product) error {
	for i, product := range products {
		if product.ID == updatedProduct.ID {
			products[i] = updatedProduct
			return nil
		}
	}

	return errors.New("Продукт для обновления не найден")
}

func DeleteProduct(products []models.Product, id models.ProductID) ([]models.Product, error) {
	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			return products, nil
		}
	}

	return products, errors.New("Продукт для удаления не найден")
}
