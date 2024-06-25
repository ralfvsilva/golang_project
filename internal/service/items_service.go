package service

import (
	"api/internal/domain"
	"time"
)

var items = []domain.Item{
	{
		ID:          1,
		Code:        "Item001",
		Title:       "Camisa",
		Description: "camisa de algodão",
		Price:       79,
		Stock:       3,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          2,
		Code:        "Item002",
		Title:       "Bola futebol",
		Description: "bola futebol",
		Price:       20,
		Stock:       30,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

type ItemService struct {
}

func (s ItemService) GetAllItems() []domain.Item {
	return items
}

func (s ItemService) AddItem(item domain.Item) []domain.Item {
	if item.Title == "" {
		return nil
	}

	items = append(items, item)

	return items
}

func (s ItemService) ReadItem(id int) *domain.Item {
	for _, item := range items {
		if id == item.ID {
			return &item
		}
	}

	return nil
}

func (s ItemService) UpdateItem(id int, itemNew domain.Item) *domain.Item {
	for i, v := range items {
		if id == v.ID {
			itemNew.ID = v.ID
			itemNew.CreatedAt = v.CreatedAt
			itemNew.UpdatedAt = time.Now()
			items[i] = itemNew

			return &itemNew
		}
	}

	return nil
}

func (s ItemService) DeleteItem(id int) error {
	var item []domain.Item
	for i, v := range item {
		if id == item[i].ID {
			item = append(item[:i], item[i+1:]...)
		}
	}
	return nil
}
