package controller

import (
	"cafe_go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAvailableItems(c *gin.Context) {
	availableItems := repository.FindAllAvailableItems()
	c.IndentedJSON(http.StatusOK, availableItems)
}
