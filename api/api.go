package api

import (
	"net/http"
	"rest-service/structs"
	"rest-service/temp_data"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOrders(o *gin.Context) {
	o.IndentedJSON(http.StatusOK, temp_data.Orders)
}

func CreateOrder(o *gin.Context) {
	var newOrder structs.Order

	if err := o.BindJSON(&newOrder); err != nil {
		return
	}

	temp_data.Orders = append(temp_data.Orders, newOrder)
	o.IndentedJSON(http.StatusCreated, newOrder)
}

func GetOrderByID(o *gin.Context) {
	id := o.Param("id")
	intVar, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	for _, a := range temp_data.Orders {

		if a.Number == intVar {
			o.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	o.IndentedJSON(http.StatusNotFound, gin.H{"message": "order not found"})
}

func GetOrderItems(o *gin.Context) {
	id := o.Param("id")
	intVar, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	ordItems := []structs.OrderItem{}

	for _, a := range temp_data.OrderItems {
		if a.OrderNumber == intVar {
			ordItems = append(ordItems, a)
		}

	}

	o.IndentedJSON(http.StatusOK, ordItems)
}
