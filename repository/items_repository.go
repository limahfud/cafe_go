package repository

import "cafe_go/model"

type ItemsRepository interface {
	FindAllAvailableItems() []model.AvailableItem
}
