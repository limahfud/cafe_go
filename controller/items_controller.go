package controller

import (
	"cafe_go/config"
	"cafe_go/repository"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAvailableItems(c *gin.Context) {
	db := config.GetDBConnection()
	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	itemsRepository := repository.NewItemsRepository()

	availableItems := itemsRepository.FindAllAvailableItems(ctx, tx)
	c.IndentedJSON(http.StatusOK, availableItems)
}
