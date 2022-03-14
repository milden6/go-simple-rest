package structs

type Order struct {
	Number        int     `json:"Number"`
	Date          string  `json:"Date"`
	ShipToParty   string  `json:"ShipToParty"`
	SoldToParty   string  `json:"SoldToParty"`
	ContactPerson string  `json:"ContactPerson"`
	Price         float64 `json:"Price"`
	Currency      string  `json:"Currency"`
}

type OrderItem struct {
	OrderNumber int    `json:"OrderNumbe"`
	ItemNumber  int    `json:"ItemNumber"`
	ItemText    string `json:"ItemText"`
	Count       int    `json:"Count"`
}
