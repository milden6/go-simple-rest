package temp_data

import "rest-service/structs"

var Orders = []structs.Order{
	{Number: 1, Date: "2022-03-01", ShipToParty: "ABN", SoldToParty: "JISK LLC", ContactPerson: "James Blunt", Price: 2000},
	{Number: 2, Date: "2022-03-02", ShipToParty: "ABN", SoldToParty: "JISK LLC", ContactPerson: "James Blunt", Price: 3000},
	{Number: 3, Date: "2022-03-03", ShipToParty: "ABN", SoldToParty: "JISK LLC", ContactPerson: "James Blunt", Price: 2789.89},
}

var OrderItems = []structs.OrderItem{
	{OrderNumber: 1, ItemNumber: 1, ItemText: "Samsung TV 32 inch", Count: 2},
	{OrderNumber: 1, ItemNumber: 2, ItemText: "Apple TV 27 inch", Count: 5},
	{OrderNumber: 1, ItemNumber: 3, ItemText: "Dell Laptop 15 inch", Count: 3},
	{OrderNumber: 1, ItemNumber: 4, ItemText: "Marshall Killburn", Count: 1},
	{OrderNumber: 2, ItemNumber: 1, ItemText: "Cofee Jacobs", Count: 2},
	{OrderNumber: 2, ItemNumber: 2, ItemText: "Coca-Cola 1L", Count: 5},
	{OrderNumber: 2, ItemNumber: 3, ItemText: "Lays", Count: 3},
	{OrderNumber: 3, ItemNumber: 1, ItemText: "Water 5L", Count: 2},
	{OrderNumber: 3, ItemNumber: 2, ItemText: "Chocolate", Count: 5},
}
