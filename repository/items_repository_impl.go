package repository

import (
	"cafe_go/model"
	"context"
	"database/sql"
)

type ItemsRepositoryImpl struct{}

func NewItemsRepository() *ItemsRepositoryImpl {
	return &ItemsRepositoryImpl{}
}

func (repo *ItemsRepositoryImpl) FindAllAvailableItems(ctx context.Context, tx *sql.Tx) []model.AvailableItem {
	query := "SELECT id, slug, name, description, price FROM items"
	rows, err := tx.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	items := []model.AvailableItem{}
	for rows.Next() {
		item := model.AvailableItem{}
		var id int

		err := rows.Scan(&id, &item.Slug, &item.Name, &item.Description, &item.Price)
		if err != nil {
			panic(err)
		}

		items = append(items, item)
	}

	return items
}
