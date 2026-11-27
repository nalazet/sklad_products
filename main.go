package main

import (
	"fmt"
	"sklad_products/models"
	"sklad_products/storage"
)

func main() {
	category1 := models.Category{
		ID:          1,
		Name:        "Молочные продукты",
		Description: "Молоко, сыр, йогурты",
	}

	category2 := models.Category{
		ID:          2,
		Name:        "Выпечка",
		Description: "Хлебобулочные и кондитерские изделия",
	}

	product1 := models.Product{
		ID:          1,
		Name:        "Молоко",
		Description: "Молоко 3.2%",
		Price:       95,
		Quantity:    20,
		CategoryID:  1,
		IsActive:    true,
	}

	product2 := models.Product{
		ID:          2,
		Name:        "Хлеб",
		Description: "Белый хлеб",
		Price:       55,
		Quantity:    15,
		CategoryID:  2,
		IsActive:    true,
	}

	var stock models.StockManager = &product1

	fmt.Println("Количество молока до поставки:", product1.Quantity)

	err := stock.AddQuantity(10)

	if err != nil {
		fmt.Println("Ошибка", err)
	}

	fmt.Println("Количество молока после поставки:", product1.Quantity)

	err = product1.RemoveQuantity(5)

	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Товар успешно списан")
		fmt.Println("Остаток молока:", product1.Quantity)
	}

	categories := []models.Category{category1, category2}
	products := []models.Product{product1, product2}

	product3 := models.Product{
		ID:          3,
		Name:        "Сыр",
		Description: "Сыр российский",
		Price:       250,
		Quantity:    12,
		CategoryID:  1,
		IsActive:    true,
	}

	products, err = storage.AddProduct(products, product3)

	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Продукт успешно добавлен")
	}

	fmt.Println("Всего видов продуктов:", len(products))

	foundProduct, err := storage.FindProductByID(products, 2)

	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Найденный продукт:", foundProduct.Name)
		fmt.Println("Цена:", foundProduct.Price, "руб.")
		fmt.Println("Количество:", foundProduct.Quantity, "шт.")
	}

	categoryMap := createCategoryMap(categories)

	showCategories(categories)
	showProducts(products, categoryMap)
}

func showCategories(categories []models.Category) {
	fmt.Println("Категории:")

	for _, category := range categories {
		fmt.Println("Название:", category.Name)
		fmt.Println("Описание:", category.Description)
	}
}

func showProducts(products []models.Product, categoryMap map[models.CategoryID]string) {
	fmt.Println("Продукты:")

	for _, product := range products {
		if product.IsActive {
			fmt.Println("Название:", product.Name)
			fmt.Println("Цена:", product.Price, "руб.")
			fmt.Println("Количество:", product.Quantity, "шт.")

			categoryName, ok := categoryMap[product.CategoryID]

			if ok {
				fmt.Println("Категория:", categoryName)
			} else {
				fmt.Println("Категория не найдена")
			}
		} else {
			fmt.Println("Товар", product.Name, "недоступен для покупки")
		}
	}
}

func createCategoryMap(categories []models.Category) map[int]string {
	categoryMap := make(map[models.CategoryID]string)

	for _, category := range categories {
		categoryMap[category.ID] = category.Name
	}
	return categoryMap
}
