package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       int
	Quantity    int
	CategoryID  int
	IsActive    bool
}

type Category struct {
	ID          int
	Name        string
	Description string
}

func main() {
	category1 := Category{
		ID:          1,
		Name:        "Молочные продукты",
		Description: "Молоко, сыр, йогурты",
	}

	category2 := Category{
		ID:          2,
		Name:        "Выпечка",
		Description: "Хлебобулочные и кондитерские изделия",
	}

	product1 := Product{
		ID:          1,
		Name:        "Молоко",
		Description: "Молоко 3.2%",
		Price:       95,
		Quantity:    20,
		CategoryID:  1,
		IsActive:    true,
	}

	product2 := Product{
		ID:          2,
		Name:        "Хлеб",
		Description: "Белый хлеб",
		Price:       55,
		Quantity:    15,
		CategoryID:  2,
		IsActive:    true,
	}

	categories := []Category{category1, category2}
	products := []Product{product1, product2}

	fmt.Println("Категории:")
	fmt.Println(categories)

	fmt.Println("Продукты:")
	fmt.Println(products)
}
