package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	setUpGin()
}

func setUpGin() {
	ginEngine := gin.Default()

	SetUpRouter(ginEngine)

	ginEngine.Run("localhost:8081")
}
