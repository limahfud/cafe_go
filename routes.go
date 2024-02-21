package main

import (
	"cafe_go/controller"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(router *gin.Engine) {
	router.GET("/available-items", controller.GetAvailableItems)
}
