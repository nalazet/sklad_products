package models

import "errors"

type ProductID = int

type Product struct {
	ID          ProductID
	Name        string
	Description string
	Price       int
	Quantity    int
	CategoryID  CategoryID
	IsActive    bool
}

type StockManager interface {
	AddQuantity(amount int) error
	RemoveQuantity(amount int) error
}

func (p *Product) AddQuantity(amount int) error {
	if amount <= 0 {
		return errors.New("Количество должно быть больше нуля")
	}

	p.Quantity = p.Quantity + amount

	return nil
}

func (p *Product) RemoveQuantity(amount int) error {
	if amount <= 0 {
		return errors.New("Количество товаров должно быть больше нуля")
	}

	if !p.IsActive {
		return errors.New("Товар недоступен для покупки")
	}

	if amount > p.Quantity {
		return errors.New("Недостаточно товара на складе")
	}

	p.Quantity = p.Quantity - amount

	return nil
}
