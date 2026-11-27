package models

type CategoryID = int

type Category struct {
	ID          CategoryID
	Name        string
	Description string
}
