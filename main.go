package main

import (
	"rest-service/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/orders", api.GetOrders)
	router.GET("/orders/:id", api.GetOrderByID)
	router.GET("/order-items/:id", api.GetOrderItems)
	router.POST("/orders", api.CreateOrder)

	router.Run("localhost:8080")
}
